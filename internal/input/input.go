package input

import (
	"io"

	"github.com/NikitaKissa/cvgen/internal/cv"
)

type Parser interface {
	Parse(r io.Reader) (cv.CV, error)
}
