package ai

import (
	"context"
	"fmt"

	"github.com/NikitaKissa/cvgen/internal/llm"
)

func (d *Dialogue) FindKeyWords(ctx context.Context, vacancy Vacancy, opts Options) (string, error) {
	const (
		systemPromptF = `
		Answer must be written in %s
		`

		userPromptF = `
		You are an ATS (Applicant Tracking System) filtering expert. 
		You are given a job description. Your task is to extract the most relevant keywords from the job description and summarize it. 
		The output should be summarization of the job description followed by a comma-separated list of keywords that are most likely to be used by an ATS to filter candidates for this position.

		Job description: 
		Position: %s
		Company: %s
		Description: 
		%s
		`
	)

	gpt, ok := llm.FromContext(ctx)
	if !ok {
		return "", fmt.Errorf("no LLM engine found in context")
	}

	request := llm.GenerateRequest{
		MaxTokens:   2048,
		Temperature: 0.5,

		Messages: []llm.Message{
			{Role: llm.RoleSystem, Content: fmt.Sprintf(systemPromptF, opts.Language)},
			{Role: llm.RoleUser, Content: fmt.Sprintf(userPromptF, vacancy.Position, vacancy.Company, vacancy.Description)},
		},
	}

	d.userRequests = append(d.userRequests, request)
	output, err := gpt.Generate(ctx, request)
	if err != nil {
		return "", fmt.Errorf("error during generating key words: %w", err)
	}

	d.aiResponses = append(d.aiResponses, output)
	return output, nil
}
