package input_json

import (
	"github.com/NikitaKissa/cvgen/internal/core"
	"github.com/NikitaKissa/cvgen/internal/cv"
)

type Certificate struct {
	Name   string         `json:"name"`
	Issuer string         `json:"issuer"`
	Url    core.StringUrl `json:"url"`
}

func (c *Certificate) ToModel() cv.Certificate {
	return cv.Certificate{
		Name:   c.Name,
		Issuer: c.Issuer,
		Url:    *c.Url.Parse(),
	}
}

func certificatesToModel(certificates []Certificate) []cv.Certificate {
	var output = make([]cv.Certificate, len(certificates))
	for i, item := range certificates {
		output[i] = item.ToModel()
	}
	return output
}
