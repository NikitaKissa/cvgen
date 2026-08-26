package generate

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/NikitaKissa/cvgen/internal/asset"
	"github.com/NikitaKissa/cvgen/internal/core"
	"github.com/NikitaKissa/cvgen/internal/cv"
	"github.com/NikitaKissa/cvgen/internal/input"
	input_json "github.com/NikitaKissa/cvgen/internal/input/json"
	"github.com/NikitaKissa/cvgen/internal/template"
)

func Run(ctx context.Context, opts Options) error {
	cv, err := parseInputFile(opts.InputPath)
	if err != nil {
		return err
	}

	if err := resolveAssets(cv); err != nil {
		return err
	}

	output, err := renderTemplate(opts, cv)
	if err != nil {
		return err
	}

	if err := writeOutputFile(opts.OutputPath, output); err != nil {
		return err
	}

	return nil
}

func parseInputFile(inputPath string) (*cv.CV, error) {
	parser, err := createParser(inputPath)
	if err != nil {
		return nil, err
	}

	file, err := os.Open(inputPath)
	if err != nil {
		return nil, fmt.Errorf(
			"%w `%s`: %s",
			core.ErrOpenFile,
			inputPath,
			err,
		)
	}
	defer file.Close()

	cv, err := parser.Parse(file)
	if err != nil {
		return nil, fmt.Errorf(
			"error during parsing input: %w",
			err,
		)
	}

	return &cv, nil
}

func createParser(inputPath string) (input.Parser, error) {
	ext := filepath.Ext(inputPath)

	switch strings.ToLower(ext) {
	case ".json":
		return input_json.New(), nil
	default:
		return nil, fmt.Errorf(
			"unknown input file extension `%s`: %w",
			ext,
			core.ErrInvalidFileExtension,
		)
	}
}

func resolveAssets(cv *cv.CV) error {
	resolvedPhoto, err := asset.LoadPhoto(
		derefOrEmpty(cv.Basics.Photo),
	)
	if err != nil {
		return fmt.Errorf(
			"resolving `basics.photo`: %w",
			err,
		)
	}

	cv.Basics.Photo = &resolvedPhoto

	return nil
}

func renderTemplate(opts Options, cv *cv.CV) ([]byte, error) {
	templateHTML, err := loadTemplateHTML(opts.TemplatePath)
	if err != nil {
		return nil, err
	}

	style, err := loadTemplateStyle(opts.StylePath)
	if err != nil {
		return nil, err
	}

	templateData := template.ModelToTemplateData(*cv, style)

	output, err := template.Render(templateHTML, templateData)
	if err != nil {
		return nil, fmt.Errorf(
			"error during rendering output html: %w",
			err,
		)
	}

	return output, nil
}

func loadTemplateHTML(templatePath string) (string, error) {
	templateHTML, err := template.LoadTemplate(templatePath)
	if err != nil {
		return "", fmt.Errorf(
			"error during loading template: %w",
			err,
		)
	}

	return templateHTML, nil
}

func loadTemplateStyle(stylePath string) (string, error) {
	style, err := template.LoadStyle(stylePath)
	if err != nil {
		return "", fmt.Errorf(
			"error during loading style: %w",
			err,
		)
	}

	return style, nil
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

func derefOrEmpty(s *string) string {
	if s == nil {
		return ""
	}

	return *s
}
