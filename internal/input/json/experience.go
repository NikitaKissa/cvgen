package input_json

import (
	"time"

	"github.com/NikitaKissa/cvgen/internal/core"
	"github.com/NikitaKissa/cvgen/internal/cv"
)

type Experience struct {
	Company        string          `json:"company"`
	CompanyWebsite *core.StringUrl `json:"company_website"`
	Position       string          `json:"position"`
	Description    *string         `json:"description"`
	Stack          []string        `json:"stack"`
	From           core.Date       `json:"from"`
	To             *core.Date      `json:"to"` // nil means "till now"
}

func (e *Experience) ToModel() cv.Experience {
	var toTime *time.Time
	if e.To != nil {
		t := e.To.ToTime()
		toTime = &t
	}

	return cv.Experience{
		Company:        e.Company,
		CompanyWebsite: e.CompanyWebsite.Parse(),
		Position:       e.Position,
		Description:    e.Description,
		Stack:          e.Stack,
		From:           e.From.ToTime(),
		To:             toTime,
	}
}

func experienceToModel(experience []Experience) []cv.Experience {
	var output = make([]cv.Experience, len(experience))
	for i, item := range experience {
		output[i] = item.ToModel()
	}
	return output
}
