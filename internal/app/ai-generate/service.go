package aigenerate

import (
	"context"
	"fmt"
	"os"

	"github.com/NikitaKissa/cvgen/internal/ai"
	"github.com/NikitaKissa/cvgen/internal/compile"
	"github.com/NikitaKissa/cvgen/internal/input"
	"github.com/NikitaKissa/cvgen/internal/template"
)

func Run(
	ctx context.Context,
	resume string,
	vacancy ai.Vacancy,
	aiOpts ai.Options,
	genOpts Options,
) error {
	dialogue := ai.NewDialogue()

	keyWords, err := dialogue.FindKeyWords(ctx, vacancy, aiOpts)
	if err != nil {
		return err
	}

	jsonCV, err := dialogue.GenerateJSON(ctx, aiOpts, keyWords, resume)
	if err != nil {
		return err
	}

	// Generation HTML from JSON
	model, err := input.LoadFromString(jsonCV)
	if err != nil {
		return fmt.Errorf("error during loading model: %w", err)
	}

	templateHTML, err := template.LoadTemplate(genOpts.TemplatePath)
	if err != nil {
		return fmt.Errorf("error during loading template: %w", err)
	}

	style, err := template.LoadStyle(genOpts.StylePath)
	if err != nil {
		return fmt.Errorf("error during loading style: %w", err)
	}

	output, err := compile.New().Compile(ctx, model, templateHTML, style)
	if err != nil {
		return fmt.Errorf("error during compiling cv: %w", err)
	}

	if err := os.WriteFile(genOpts.OutputPath, output, 0644); err != nil {
		return fmt.Errorf("error during writing output: %w", err)
	}

	return nil
}
