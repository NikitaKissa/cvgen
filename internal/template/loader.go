package template

import (
	_ "embed"
	"fmt"
	"os"
)

//go:embed defaults/template.html
var DefaultTemplateHTML string

//go:embed defaults/style.css
var DefaultStyleCSS string

// LoadTemplate returns the HTML template at path, or the embedded
// default template when path is empty.
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

// LoadStyle returns the CSS at path, or the embedded default CSS when
// path is empty.
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
