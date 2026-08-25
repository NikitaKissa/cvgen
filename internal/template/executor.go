package template

import (
	"bytes"
	"fmt"
	"html/template"
)

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
