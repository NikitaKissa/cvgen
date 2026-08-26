package verify

import (
	"context"
	"fmt"

	"github.com/NikitaKissa/cvgen/internal/compile"
	"github.com/NikitaKissa/cvgen/internal/input"
	"github.com/NikitaKissa/cvgen/internal/template"
)

// Report summarizes a successful verification. It is informational only:
// its presence means input, template and style all parsed and compiled
// without errors.
type Report struct {
	Experience   int
	Education    int
	SkillGroups  int
	Certificates int
	Projects     int
	Languages    int
	Clause       bool

	// OutputSize is the size in bytes of the HTML that *would* be
	// produced. Nothing is written to disk.
	OutputSize int
}

// Run parses the CV JSON, loads the template and style (falling back to
// embedded defaults exactly like `generate` does), and runs them through
// the same compiler `generate` uses — without ever writing a file.
//
// Any error returned here is exactly the error `generate` would have
// hit at the equivalent stage, so a clean Run here guarantees `generate`
// will succeed with the same inputs.
func Run(ctx context.Context, opts Options) (Report, error) {
	model, err := input.Load(opts.InputPath)
	if err != nil {
		return Report{}, fmt.Errorf("invalid input: %w", err)
	}

	templateHTML, err := template.LoadTemplate(opts.TemplatePath)
	if err != nil {
		return Report{}, fmt.Errorf("invalid template: %w", err)
	}

	style, err := template.LoadStyle(opts.StylePath)
	if err != nil {
		return Report{}, fmt.Errorf("invalid style: %w", err)
	}

	output, err := compile.New().Compile(ctx, model, templateHTML, style)
	if err != nil {
		return Report{}, fmt.Errorf("compilation failed: %w", err)
	}

	return Report{
		Experience:   len(model.Experience),
		Education:    len(model.Education),
		SkillGroups:  len(model.SkillGroups),
		Certificates: len(model.Certificates),
		Projects:     len(model.Projects),
		Languages:    len(model.Languages),
		Clause:       model.Clause != nil,
		OutputSize:   len(output),
	}, nil
}
