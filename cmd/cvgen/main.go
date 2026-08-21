package main

import (
	"fmt"
	"os"

	"github.com/NikitaKissa/cvgen/internal/input"
	input_json "github.com/NikitaKissa/cvgen/internal/input/json"
	"github.com/NikitaKissa/cvgen/internal/template"
)

func main() {
	// template import to binary
	_ = template.DefaultTemplateHTML
	_ = template.DefaultStyleCSS
	// ^^^ DON'T touch this part ^^^

	// File processing section

	file, err := os.Open("./testdata/valid/cv.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Model section

	var parser input.Parser = input_json.New()

	cv, err := parser.Parse(file)
	if err != nil {
		panic(err)
	}

	// Template section

	templateHtml, err := template.LoadTemplate("")
	if err != nil {
		err = fmt.Errorf("error during loading template: %w", err)
		panic(err)
	}
	style, err := template.LoadStyle("")
	if err != nil {
		err = fmt.Errorf("error during loading style: %w", err)
		panic(err)
	}

	templateData := template.ModelToTemplateData(cv)
	output, err := template.Render(templateHtml, style, templateData)
	if err != nil {
		panic(err)
	}

	// Output

	err = os.WriteFile("/tmp/test-cv.html", output, 0664)
	if err != nil {
		panic(err)
	}
}
