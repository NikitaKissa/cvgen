package template

import (
	"html/template"
	"time"

	"github.com/NikitaKissa/cvgen/internal/cv"
)

type TemplateData struct {
	Style        template.CSS
	Basics       Basics
	Education    []Education
	Experience   []Experience
	SkillGroups  []cv.SkillGroup
	Certificates []Certificate
	Projects     []Project
	Languages    []cv.Language
	Clause       *string
}

type Basics struct {
	FullName    string
	Position    string
	Description string
	Email       *string
	Phone       *string
	Photo       template.URL
	Links       Links
}

type Links map[string]template.URL

type Education struct {
	Institution        string
	InstitutionWebsite *template.URL
	FieldOfStudy       string
	Description        *string
	From               time.Time
	To                 *time.Time
}

type Experience struct {
	Company        string
	CompanyWebsite *template.URL
	Position       string
	Description    *string
	Stack          []string
	From           time.Time
	To             *time.Time // nil means "till now"
}

type Certificate struct {
	Name   string
	Issuer string
	Url    template.URL
}

type Project struct {
	Name        string
	Description *string
	Url         template.URL
}
