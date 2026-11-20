package llama

/*
#cgo CFLAGS: -I/home/nikita/projects/cvgen/native/linux-amd64/include
#cgo CFLAGS: -I/home/nikita/projects/cvgen/native/linux-amd64/include/ggml
#cgo LDFLAGS: -L/home/nikita/projects/cvgen/native/linux-amd64/lib -lllama

#include <stdlib.h>
#include "llama.h"
*/
import "C"

import (
	"context"
	"fmt"
	"math"
	"unsafe"

	"github.com/NikitaKissa/cvgen/internal/llm"
)

const maxGeneratedTokens = 512

func loadModel(path string) (llamaModel, error) {
	cPath := C.CString(path)
	defer C.free(unsafe.Pointer(cPath))

	params := C.llama_model_default_params()

	model := C.llama_model_load_from_file(cPath, params)
	if model == nil {
		return llamaModel{}, fmt.Errorf("failed to load model: %s", path)
	}

	return llamaModel{
		ptr: uintptr(unsafe.Pointer(model)),
	}, nil
}

func createContext(model llamaModel) (llamaContext, error) {
	if model.ptr == 0 {
		return llamaContext{}, fmt.Errorf("llama: invalid model")
	}

	params := C.llama_context_default_params()

	ctx := C.llama_init_from_model(
		(*C.struct_llama_model)(unsafe.Pointer(model.ptr)),
		params,
	)

	if ctx == nil {
		return llamaContext{}, fmt.Errorf("failed to create llama context")
	}

	return llamaContext{
		ptr: uintptr(unsafe.Pointer(ctx)),
	}, nil
}

func createSampler() (llamaSampler, error) {
	params := C.llama_sampler_chain_default_params()

	sampler := C.llama_sampler_chain_init(params)
	if sampler == nil {
		return llamaSampler{}, fmt.Errorf("failed to create sampler chain")
	}

	// The chain takes ownership of the sampler.
	C.llama_sampler_chain_add(
		sampler,
		C.llama_sampler_init_greedy(),
	)

	return llamaSampler{
		ptr: uintptr(unsafe.Pointer(sampler)),
	}, nil
}

func freeModel(model llamaModel) {
	if model.ptr == 0 {
		return
	}

	C.llama_model_free(
		(*C.struct_llama_model)(unsafe.Pointer(model.ptr)),
	)
}

func freeContext(ctx llamaContext) {
	if ctx.ptr == 0 {
		return
	}

	C.llama_free(
		(*C.struct_llama_context)(unsafe.Pointer(ctx.ptr)),
	)
}

func freeSampler(sampler llamaSampler) {
	if sampler.ptr == 0 {
		return
	}

	C.llama_sampler_free(
		(*C.struct_llama_sampler)(unsafe.Pointer(sampler.ptr)),
	)
}

func generate(
	ctx context.Context,
	model llamaModel,
	llamaCtx llamaContext,
	sampler llamaSampler,
	req llm.GenerateRequest,
) (string, error) {
	if req.Prompt == "" {
		return "", fmt.Errorf("llama: prompt is empty")
	}

	if model.ptr == 0 || llamaCtx.ptr == 0 || sampler.ptr == 0 {
		return "", fmt.Errorf("llama: engine is not initialized")
	}

	if len(req.Prompt) > math.MaxInt32 {
		return "", fmt.Errorf("llama: prompt is too large")
	}

	if req.MaxTokens <= 0 {
		req.MaxTokens = maxGeneratedTokens
	}

	cModel := (*C.struct_llama_model)(unsafe.Pointer(model.ptr))
	cCtx := (*C.struct_llama_context)(unsafe.Pointer(llamaCtx.ptr))
	cSampler := (*C.struct_llama_sampler)(unsafe.Pointer(sampler.ptr))

	vocab := C.llama_model_get_vocab(cModel)
	if vocab == nil {
		return "", fmt.Errorf("llama: failed to get model vocabulary")
	}

	// Tokenize prompt.
	cPrompt := C.CString(req.Prompt)
	defer C.free(unsafe.Pointer(cPrompt))

	tokenCount := C.llama_tokenize(
		vocab,
		cPrompt,
		C.int32_t(len(req.Prompt)),
		nil,
		0,
		true,
		false,
	)

	if tokenCount == -2147483648 {
		return "", fmt.Errorf("llama: prompt tokenization overflow")
	}

	if tokenCount < 0 {
		tokenCount = -tokenCount
	}

	if tokenCount == 0 {
		return "", fmt.Errorf("llama: prompt produced no tokens")
	}

	tokens := make([]C.llama_token, int(tokenCount))

	tokenCount = C.llama_tokenize(
		vocab,
		cPrompt,
		C.int32_t(len(req.Prompt)),
		&tokens[0],
		C.int32_t(len(tokens)),
		true,
		false,
	)

	if tokenCount < 0 {
		return "", fmt.Errorf("llama: failed to tokenize prompt")
	}

	// Create batch.
	batch := C.llama_batch_ext_init(cCtx)
	if batch == nil {
		return "", fmt.Errorf("llama: failed to create batch")
	}
	defer C.llama_batch_ext_free(batch)

	// Add prompt tokens to the batch.
	for i, token := range tokens {
		index := C.llama_batch_ext_add_token(
			batch,
			0,
			token,
		)

		if index < 0 {
			return "", fmt.Errorf("llama: failed to add prompt token: %d", index)
		}

		position := C.llama_pos(i)

		if !bool(C.llama_batch_ext_set_pos(batch, index, &position)) {
			return "", fmt.Errorf("llama: failed to set prompt token position")
		}

		// We only need logits for the final prompt token.
		if i == len(tokens)-1 {
			if !bool(C.llama_batch_ext_set_output_logits(batch, index, true)) {
				return "", fmt.Errorf("llama: failed to enable prompt logits")
			}
		}
	}

	// Process prompt.
	result := C.llama_process(
		cCtx,
		C.LLAMA_PROCESS_TYPE_DECODE,
		batch,
	)

	if result != 0 {
		return "", fmt.Errorf("llama: prompt processing failed: %d", result)
	}

	output := make([]byte, 0)

	// Generate tokens.
	for i := 0; i < req.MaxTokens; i++ {
		if err := ctx.Err(); err != nil {
			return string(output), err
		}

		token := C.llama_sampler_sample(cSampler, cCtx, -1)

		if bool(C.llama_vocab_is_eog(vocab, token)) {
			break
		}

		piece, err := tokenToPiece(vocab, token)
		if err != nil {
			return string(output), err
		}

		output = append(output, piece...)

		// Prepare batch for the next token.
		C.llama_batch_ext_clear(batch)

		index := C.llama_batch_ext_add_token(
			batch,
			0,
			token,
		)

		if index < 0 {
			return string(output), fmt.Errorf("llama: failed to add generated token: %d", index)
		}

		position := C.llama_pos(len(tokens) + i)

		if !bool(C.llama_batch_ext_set_pos(batch, index, &position)) {
			return string(output), fmt.Errorf("llama: failed to set generated token position")
		}

		if !bool(C.llama_batch_ext_set_output_logits(batch, index, true)) {
			return string(output), fmt.Errorf("llama: failed to enable generated token logits")
		}

		result := C.llama_process(
			cCtx,
			C.LLAMA_PROCESS_TYPE_DECODE,
			batch,
		)

		if result != 0 {
			return string(output), fmt.Errorf("llama: token processing failed: %d", result)
		}
	}

	return string(output), nil
}

func tokenToPiece(
	vocab *C.struct_llama_vocab,
	token C.llama_token,
) ([]byte, error) {
	buffer := make([]C.char, 256)

	n := C.llama_token_to_piece(
		vocab,
		token,
		&buffer[0],
		C.int32_t(len(buffer)),
		0,
		false,
	)

	// Negative value means the buffer was too small.
	if n < 0 {
		size := -n
		buffer = make([]C.char, int(size))

		n = C.llama_token_to_piece(
			vocab,
			token,
			&buffer[0],
			C.int32_t(len(buffer)),
			0,
			false,
		)
	}

	if n < 0 {
		return nil, fmt.Errorf("llama: failed to convert token to text")
	}

	result := make([]byte, int(n))
	for i := 0; i < int(n); i++ {
		result[i] = byte(buffer[i])
	}

	return result, nil
}
