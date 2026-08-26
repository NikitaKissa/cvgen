package input

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/NikitaKissa/cvgen/internal/core"
	"github.com/NikitaKissa/cvgen/internal/cv"
	input_json "github.com/NikitaKissa/cvgen/internal/input/json"
)

// Load opens the file at path, picks a Parser based on its extension,
// and parses it into a CV domain model.
//
// This is where input format dispatch lives: callers (app/generate,
// app/verify, ...) only ever deal with cv.CV, never with "is this
// JSON/YAML/TOML".
func Load(path string) (cv.CV, error) {
	parser, err := newParser(path)
	if err != nil {
		return cv.CV{}, err
	}

	file, err := os.Open(path)
	if err != nil {
		return cv.CV{}, fmt.Errorf(
			"%w `%s`: %s",
			core.ErrOpenFile,
			path,
			err,
		)
	}
	defer file.Close()

	model, err := parser.Parse(file)
	if err != nil {
		return cv.CV{}, fmt.Errorf("%w: %s", ErrParsing, err)
	}

	return model, nil
}

func newParser(path string) (Parser, error) {
	ext := filepath.Ext(path)

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
