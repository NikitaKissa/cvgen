package input_json

import (
	"net/url"

	"github.com/NikitaKissa/cvgen/internal/cv"
)

type Certificate struct {
	Name   string  `json:"name"`
	Issuer string  `json:"issuer"`
	Url    url.URL `json:"url"`
}

func (c *Certificate) ToModel() cv.Certificate {
	return cv.Certificate{
		Name:   c.Name,
		Issuer: c.Issuer,
		Url:    c.Url,
	}
}

func certificatesToModel(certificates []Certificate) []cv.Certificate {
	var output = make([]cv.Certificate, len(certificates))
	for i, item := range certificates {
		output[i] = item.ToModel()
	}
	return output
}
