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
	var inputData InputDTO

	dec := json.NewDecoder(r)
	if err := dec.Decode(&inputData); err != nil {
		return cv.CV{}, fmt.Errorf(
			"error during decoding json: %w",
			err,
		)
	}

	// Validation is for future
	// if err := inputData.Validate(); err != nil {
	// 	return cv.CV{}, fmt.Errorf("json validation error: %w", err)
	// }

	result := inputData.ToModel()
	return result, nil
}
