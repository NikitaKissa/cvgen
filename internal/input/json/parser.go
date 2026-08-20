package input_json

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/NikitaKissa/cvgen/internal/cv"
)

type Parser struct{}

func New() *Parser {
	return &Parser{}
}

func (p *Parser) Parse(r io.Reader) (cv.CV, error) {
	var input InputDTO

	dec := json.NewDecoder(r)
	if err := dec.Decode(&input); err != nil {
		return cv.CV{}, fmt.Errorf("error during decoding json: %w", err)
	}

	// Validation is for future
	// if err := input.Validate(); err != nil {
	// 	return cv.CV{}, fmt.Errorf("json validation error: %w", err)
	// }

	result := input.ToModel()
	return result, nil
}
