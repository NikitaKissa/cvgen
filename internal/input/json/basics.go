package input_json

import (
	"net/url"

	"github.com/NikitaKissa/cvgen/internal/core"
	"github.com/NikitaKissa/cvgen/internal/cv"
)

type Basics struct {
	Photo       *string `json:"photo"`
	FullName    string  `json:"full_name"`
	Position    string  `json:"position"`
	Description string  `json:"description"`
	Email       *string `json:"email"`
	Phone       *string `json:"phone"`
	Links       Links   `json:"links"`
}

type Links map[string]core.StringUrl // ex: "LinkedIn": "https://www.linkedin.com/in/mykyta-kissa-684aa833a"

func (b *Basics) ToModel() cv.Basics {
	return cv.Basics{
		Photo:       b.Photo,
		FullName:    b.FullName,
		Position:    b.Position,
		Description: b.Description,
		Email:       b.Email,
		Phone:       b.Phone,
		Links:       stringToLinks(b.Links),
	}
}

func stringToLinks(strings Links) cv.Links {
	links := make(map[string]url.URL, len(strings))

	for name, rawURL := range strings {
		u := rawURL.Parse()
		links[name] = *u
	}
	return links
}
