package generate

import (
	"context"
	"fmt"
	"io"
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

	if err := writeOutput(opts.Stdout, opts.OutputPath, output); err != nil {
		return err
	}

	return nil
}

func writeOutput(isStdout bool, outputPath string, output []byte) error {
	var w io.Writer
	if isStdout {
		w = os.Stdout
	} else {
		file, err := os.Create(outputPath)
		if err != nil {
			return fmt.Errorf(
				"%w `%s`: %s",
				core.ErrWriteFile,
				outputPath,
				err,
			)
		}
		defer file.Close()

		w = file
	}

	if _, err := w.Write(output); err != nil {
		return fmt.Errorf("%w: %s", core.ErrWriteOutput, err)
	}

	return nil
}
