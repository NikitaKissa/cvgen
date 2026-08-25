package template

import (
	"html/template"

	"github.com/NikitaKissa/cvgen/internal/cv"
)

type TemplateData struct {
	Style        template.CSS
	Basics       Basics
	Education    []cv.Education
	Experience   []cv.Experience
	SkillGroups  []cv.SkillGroup
	Certificates []cv.Certificate
	Projects     []cv.Project
	Languages    []cv.Language
}

type Basics struct {
	FullName    string
	Position    string
	Description string
	Email       *string
	Phone       *string
	Photo       template.URL
	Links       cv.Links
}
