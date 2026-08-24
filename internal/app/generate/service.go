package generate

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/NikitaKissa/cvgen/internal/core"
	"github.com/NikitaKissa/cvgen/internal/input"
	input_json "github.com/NikitaKissa/cvgen/internal/input/json"
	"github.com/NikitaKissa/cvgen/internal/template"
)

func Run(ctx context.Context, opts Options) error {
	// Extension check section
	ext := filepath.Ext(opts.InputPath)
	var parser input.Parser
	switch strings.ToLower(ext) {
	case ".json":
		parser = input_json.New()
	default:
		return fmt.Errorf(
			"unknown input file extension `%s`: %w",
			ext,
			core.ErrInvalidFileExtension,
		)
	}

	file, err := os.Open(opts.InputPath)
	if err != nil {
		return fmt.Errorf(
			"%w `%s`: %s",
			core.ErrOpenFile,
			opts.InputPath,
			err,
		)
	}
	defer file.Close()

	// Model section

	cv, err := parser.Parse(file)
	if err != nil {
		return fmt.Errorf(
			"error during parsing input: %w",
			err,
		)
	}
	file.Close()

	// Template section

	templateHtml, err := template.LoadTemplate(opts.TemplatePath)
	if err != nil {
		return fmt.Errorf("error during loading template: %w", err)

	}
	style, err := template.LoadStyle(opts.StylePath)
	if err != nil {
		return fmt.Errorf("error during loading style: %w", err)
	}

	templateData := template.ModelToTemplateData(cv, style)
	output, err := template.Render(templateHtml, templateData)
	if err != nil {
		return fmt.Errorf("error during rendering output html: %w", err)
	}

	// Output

	err = os.WriteFile(opts.OutputPath, output, 0664)
	if err != nil {
		return fmt.Errorf(
			"%w `%s`: %s",
			core.ErrWriteFile,
			opts.OutputPath,
			err,
		)
	}

	return nil
}
