package template

import (
	"bytes"
	"fmt"
	"html/template"

	"github.com/NikitaKissa/cvgen/internal/cv"
)

type TemplateData struct {
	Style        template.CSS
	Basics       cv.Basics
	Education    []cv.Education
	Experience   []cv.Experience
	SkillGroups  []cv.SkillGroup
	Certificates []cv.Certificate
	Projects     []cv.Project
	Languages    []cv.Language
}

func Render(templateHtml string, data TemplateData) ([]byte, error) {
	tmpl, err := template.New("cv").Funcs(FuncMap()).Parse(templateHtml)
	if err != nil {
		return nil, fmt.Errorf("parse template: %w", err)
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return nil, fmt.Errorf("execute template: %w", err)
	}

	return buf.Bytes(), nil
}
