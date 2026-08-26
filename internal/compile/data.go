package compile

import (
	"html/template"

	"github.com/NikitaKissa/cvgen/internal/cv"
	tmpl "github.com/NikitaKissa/cvgen/internal/template"
)

// ModelToTemplateData converts the domain model plus a resolved style
// string into the presentation-facing structure that user templates see.
//
// This is the boundary between "what the CV is" (cv.CV) and "what the
// template gets" (tmpl.TemplateData) — it belongs here, in compile, and
// not in internal/template, because internal/template must stay
// unaware of the CV domain model.
func ModelToTemplateData(model cv.CV, style string) tmpl.TemplateData {
	return tmpl.TemplateData{
		Style: template.CSS(style),
		Basics: tmpl.Basics{
			FullName:    model.Basics.FullName,
			Position:    model.Basics.Position,
			Description: model.Basics.Description,
			Email:       model.Basics.Email,
			Phone:       model.Basics.Phone,
			Photo:       template.URL(derefOrEmpty(model.Basics.Photo)),
			Links:       model.Basics.Links,
		},
		Education:    model.Education,
		Experience:   model.Experience,
		SkillGroups:  model.SkillGroups,
		Certificates: model.Certificates,
		Projects:     model.Projects,
		Languages:    model.Languages,
		Clause:       model.Clause,
	}
}

func derefOrEmpty(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}
