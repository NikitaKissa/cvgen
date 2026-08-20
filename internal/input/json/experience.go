package input_json

import (
	"net/url"
	"time"

	"github.com/NikitaKissa/cvgen/internal/cv"
)

type Experience struct {
	Company        string     `json:"company"`
	CompanyWebsite *url.URL   `json:"company_website"`
	Position       string     `json:"position"`
	Description    *string    `json:"description"`
	Stack          []string   `json:"stack"`
	From           time.Time  `json:"from"`
	To             *time.Time `json:"to"` // nil means "till now"
}

func (e *Experience) ToModel() cv.Experience {
	return cv.Experience{
		Company:        e.Company,
		CompanyWebsite: e.CompanyWebsite,
		Position:       e.Position,
		Description:    e.Description,
		Stack:          e.Stack,
		From:           e.From,
		To:             e.To,
	}
}

func experienceToModel(experience []Experience) []cv.Experience {
	var output = make([]cv.Experience, len(experience))
	for i, item := range experience {
		output[i] = item.ToModel()
	}
	return output
}
