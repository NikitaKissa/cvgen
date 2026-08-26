package compile

import (
	"fmt"

	"github.com/NikitaKissa/cvgen/internal/asset"
	"github.com/NikitaKissa/cvgen/internal/cv"
)

// resolveAssets embeds optional external assets (currently: the photo)
// directly into the CV model as data URIs (or leaves remote URLs as-is),
// so the HTML produced by the compiler stays self-contained.
//
// It returns a copy of the model with resolved assets; it never mutates
// the caller's CV.
func resolveAssets(model cv.CV) (cv.CV, error) {
	resolvedPhoto, err := asset.LoadPhoto(derefOrEmpty(model.Basics.Photo))
	if err != nil {
		return cv.CV{}, fmt.Errorf("resolving `basics.photo`: %w", err)
	}

	model.Basics.Photo = &resolvedPhoto

	return model, nil
}
