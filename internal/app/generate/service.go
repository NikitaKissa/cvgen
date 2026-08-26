package generate

import (
	"context"
	"fmt"
	"os"

	"github.com/NikitaKissa/cvgen/internal/compile"
	"github.com/NikitaKissa/cvgen/internal/core"
	"github.com/NikitaKissa/cvgen/internal/input"
	"github.com/NikitaKissa/cvgen/internal/template"
)

func Run(ctx context.Context, opts Options) error {
	model, err := input.Load(opts.InputPath)
	if err != nil {
		return fmt.Errorf("error during parsing input: %w", err)
	}

	templateHTML, err := template.LoadTemplate(opts.TemplatePath)
	if err != nil {
		return fmt.Errorf("error during loading template: %w", err)
	}

	style, err := template.LoadStyle(opts.StylePath)
	if err != nil {
		return fmt.Errorf("error during loading style: %w", err)
	}

	output, err := compile.New().Compile(ctx, model, templateHTML, style)
	if err != nil {
		return fmt.Errorf("error during compiling cv: %w", err)
	}

	if err := writeOutputFile(opts.OutputPath, output); err != nil {
		return err
	}

	return nil
}

func writeOutputFile(outputPath string, output []byte) error {
	if err := os.WriteFile(outputPath, output, 0664); err != nil {
		return fmt.Errorf(
			"%w `%s`: %s",
			core.ErrWriteFile,
			outputPath,
			err,
		)
	}

	return nil
}
