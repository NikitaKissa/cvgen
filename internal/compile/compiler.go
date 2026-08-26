package compile

import (
	"context"
	"fmt"

	"github.com/NikitaKissa/cvgen/internal/cv"
	"github.com/NikitaKissa/cvgen/internal/template"
)

// Compiler combines a CV domain model, a Go HTML template and CSS into a
// single self-contained HTML document.
//
// It is the only place in the project that knows how to put
// CV + template + style + assets together. It does not know about the
// CLI, about JSON, or about the filesystem beyond what internal/asset
// already resolves for it.
type Compiler struct{}

func New() *Compiler {
	return &Compiler{}
}

// Compile renders templateHTML against model, with style inlined and
// assets (e.g. the photo) embedded as data URIs.
func (c *Compiler) Compile(
	ctx context.Context,
	model cv.CV,
	templateHTML string,
	style string,
) ([]byte, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	resolved, err := resolveAssets(model)
	if err != nil {
		return nil, fmt.Errorf("resolving assets: %w", err)
	}

	data := ModelToTemplateData(resolved, style)

	output, err := template.Render(templateHTML, data)
	if err != nil {
		return nil, fmt.Errorf("rendering template: %w", err)
	}

	return output, nil
}
