package ai

import (
	"context"
	_ "embed"
	"fmt"

	"github.com/NikitaKissa/cvgen/internal/llm"
)

// TODO: normally implement json schema for ai

//go:embed../../testdata/valid/cv.json
var exampleJson string

func (d *Dialogue) GenerateJSON(ctx context.Context, opts Options, keyWords string, resume string) (string, error) {
	const (
		systemPromptF = `
		Answer must be written in %s
		`

		userPromptF = `
		You are an ATS (Applicant Tracking System) filtering expert. 
		You are given a job description with keywords which more likely to be used by an ATS and a resume (in JSON format). 
		Adapt the resume to the job description by adding the keywords to the resume in a natural way. Also adapt the 'basic.description' field of the resume to the job description.
		The output should be ONLY a JSON object representing the resume adapted to the job description and keywords.

		JSON schema of the resume:
		%s

		Job description and keywords:
		%s

		Resume:
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
			{Role: llm.RoleUser, Content: fmt.Sprintf(userPromptF, exampleJson, keyWords, resume)},
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
