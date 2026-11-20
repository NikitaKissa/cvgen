package ai

import "github.com/NikitaKissa/cvgen/internal/llm"

type Vacancy struct {
	Position    string
	Company     string
	Description string
}

type Options struct {
	Language string
}

type Dialogue struct {
	userRequests []llm.GenerateRequest
	aiResponses  []string
}

func NewDialogue() *Dialogue {
	return &Dialogue{}
}
