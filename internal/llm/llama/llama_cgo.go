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

func createSampler(temperature float32) (llamaSampler, error) {
	params := C.llama_sampler_chain_default_params()

	sampler := C.llama_sampler_chain_init(params)
	if sampler == nil {
		return llamaSampler{}, fmt.Errorf("failed to create sampler chain")
	}

	if temperature <= 0 {
		C.llama_sampler_chain_add(sampler, C.llama_sampler_init_greedy())
	} else {
		C.llama_sampler_chain_add(sampler, C.llama_sampler_init_temp(C.float(temperature)))
		C.llama_sampler_chain_add(sampler, C.llama_sampler_init_dist(C.LLAMA_DEFAULT_SEED))
	}

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
	chatTemplate string,
	rawPrompt bool,
	req llm.GenerateRequest,
) (string, error) {
	if model.ptr == 0 || llamaCtx.ptr == 0 {
		return "", fmt.Errorf("llama: engine is not initialized")
	}

	if req.Prompt != "" && len(req.Messages) > 0 {
		return "", fmt.Errorf("llama: request cannot contain both prompt and messages")
	}
	if rawPrompt && len(req.Messages) > 0 {
		return "", fmt.Errorf("llama: messages are not supported in raw prompt mode")
	}
	if req.Prompt == "" && len(req.Messages) == 0 {
		return "", fmt.Errorf("llama: prompt is empty")
	}

	prompt := req.Prompt
	if !rawPrompt {
		messages := req.Messages
		if len(messages) == 0 {
			messages = []llm.Message{{Role: "user", Content: req.Prompt}}
		}

		var err error
		prompt, err = applyChatTemplate(model, chatTemplate, messages)
		if err != nil {
			return "", err
		}
	}

	if len(prompt) > math.MaxInt32 {
		return "", fmt.Errorf("llama: prompt is too large")
	}

	if req.MaxTokens <= 0 {
		req.MaxTokens = maxGeneratedTokens
	}
	parseSpecial := !rawPrompt

	cModel := (*C.struct_llama_model)(unsafe.Pointer(model.ptr))
	cCtx := (*C.struct_llama_context)(unsafe.Pointer(llamaCtx.ptr))
	sampler, err := createSampler(req.Temperature)
	if err != nil {
		return "", err
	}
	defer freeSampler(sampler)
	cSampler := (*C.struct_llama_sampler)(unsafe.Pointer(sampler.ptr))

	vocab := C.llama_model_get_vocab(cModel)
	if vocab == nil {
		return "", fmt.Errorf("llama: failed to get model vocabulary")
	}

	// Tokenize prompt.
	cPrompt := C.CString(prompt)
	defer C.free(unsafe.Pointer(cPrompt))

	tokenCount := C.llama_tokenize(
		vocab,
		cPrompt,
		C.int32_t(len(prompt)),
		nil,
		0,
		true,
		C.bool(parseSpecial),
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
		C.int32_t(len(prompt)),
		&tokens[0],
		C.int32_t(len(tokens)),
		true,
		C.bool(parseSpecial),
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

func applyChatTemplate(model llamaModel, template string, messages []llm.Message) (string, error) {
	cModel := (*C.struct_llama_model)(unsafe.Pointer(model.ptr))
	var cTemplate *C.char
	if template != "" {
		cTemplate = C.CString(template)
		defer C.free(unsafe.Pointer(cTemplate))
	} else {
		cTemplate = (*C.char)(unsafe.Pointer(C.llama_model_chat_template(cModel, nil)))
		if cTemplate == nil {
			return "", fmt.Errorf("llama: model has no chat template; set one with llama.WithChatTemplate or enable raw prompt mode")
		}
	}

	chatMemory := C.malloc(C.size_t(len(messages)) * C.size_t(unsafe.Sizeof(C.struct_llama_chat_message{})))
	if chatMemory == nil {
		return "", fmt.Errorf("llama: failed to allocate chat messages")
	}
	defer C.free(chatMemory)
	chat := unsafe.Slice((*C.struct_llama_chat_message)(chatMemory), len(messages))
	var allocatedStrings []unsafe.Pointer
	defer func() {
		for _, pointer := range allocatedStrings {
			C.free(pointer)
		}
	}()

	for i, message := range messages {
		if message.Role == "" {
			return "", fmt.Errorf("llama: message role is empty")
		}
		role := C.CString(message.Role)
		content := C.CString(message.Content)
		allocatedStrings = append(allocatedStrings, unsafe.Pointer(role), unsafe.Pointer(content))
		chat[i].role = role
		chat[i].content = content
	}

	length := C.llama_chat_apply_template(cTemplate, (*C.struct_llama_chat_message)(chatMemory), C.size_t(len(messages)), true, nil, 0)
	if length <= 0 {
		return "", fmt.Errorf("llama: failed to apply chat template")
	}
	if length == math.MaxInt32 {
		return "", fmt.Errorf("llama: formatted prompt is too large")
	}

	buffer := make([]C.char, int(length)+1)
	formattedLength := C.llama_chat_apply_template(
		cTemplate,
		(*C.struct_llama_chat_message)(chatMemory),
		C.size_t(len(messages)),
		true,
		&buffer[0],
		C.int32_t(len(buffer)),
	)
	if formattedLength < 0 || formattedLength > length {
		return "", fmt.Errorf("llama: failed to format chat prompt")
	}

	return C.GoStringN(&buffer[0], C.int(formattedLength)), nil
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
