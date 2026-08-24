package template

import (
	_ "embed"
	"fmt"
	"html/template"
	"os"

	"github.com/NikitaKissa/cvgen/internal/cv"
)

//go:embed defaults/template.html
var DefaultTemplateHTML string

//go:embed defaults/style.css
var DefaultStyleCSS string

func LoadTemplate(path string) (string, error) {
	if path == "" {
		return DefaultTemplateHTML, nil
	}
	data, err := os.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("read template: %w", err)
	}
	return string(data), nil
}

func LoadStyle(path string) (string, error) {
	if path == "" {
		return DefaultStyleCSS, nil
	}
	data, err := os.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("read style: %w", err)
	}
	return string(data), nil
}

func ModelToTemplateData(cv cv.CV, style string) TemplateData {
	return TemplateData{
		Style:        template.CSS(style),
		Basics:       cv.Basics,
		Education:    cv.Education,
		Experience:   cv.Experience,
		SkillGroups:  cv.SkillGroups,
		Certificates: cv.Certificates,
		Projects:     cv.Projects,
		Languages:    cv.Languages,
	}
}
