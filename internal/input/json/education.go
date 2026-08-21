package input_json

import (
	"time"

	"github.com/NikitaKissa/cvgen/internal/core"
	"github.com/NikitaKissa/cvgen/internal/cv"
)

type Education struct {
	Institution        string          `json:"institution"`
	InstitutionWebsite *core.StringUrl `json:"institution_website"`
	FieldOfStudy       string          `json:"field_of_study"`
	Description        *string         `json:"description"`
	From               time.Time       `json:"from"`
	To                 *time.Time      `json:"to"` // nil means "till now"
}

func (e *Education) ToModel() cv.Education {
	return cv.Education{
		Institution:        e.Institution,
		InstitutionWebsite: e.InstitutionWebsite.Parse(),
		FieldOfStudy:       e.FieldOfStudy,
		Description:        e.Description,
		From:               e.From,
		To:                 e.To,
	}
}

func educationToModel(education []Education) []cv.Education {
	var output = make([]cv.Education, len(education))
	for i, item := range education {
		output[i] = item.ToModel()
	}
	return output
}
