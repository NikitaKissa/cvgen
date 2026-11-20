package main

import (
	"context"
	"fmt"
	"time"

	"github.com/NikitaKissa/cvgen/internal/llm"
	"github.com/NikitaKissa/cvgen/internal/llm/llama"
)

func main() {
	ctx := context.Background()
	gpt := llama.New()
	defer gpt.Close()

	fmt.Println("Loading model...")
	err := gpt.LoadModel(ctx, "/home/nikita/downloads/chrome/Qwen3.5-4B-Q4_K_M.gguf")
	if err != nil {
		fmt.Println("Error loading model:", err)
		return
	}
	fmt.Println("Model loaded.")

	start := time.Now()
	a, _ := gpt.Generate(ctx, llm.DefaultRequest("2 + 2 = ?"))

	fmt.Println(a)
	fmt.Println("Time taken:", time.Since(start))
}
