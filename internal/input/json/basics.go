package input_json

import (
	"net/url"

	"github.com/NikitaKissa/cvgen/internal/cv"
)

type Basics struct {
	FullName    string  `json:"full_name"`
	Position    string  `json:"position"`
	Description string  `json:"description"`
	Email       *string `json:"email"`
	Phone       *string `json:"phone"`
	Links       Links   `json:"links"`
}

type Links map[string]url.URL // ex: "LinkedIn": "https://www.linkedin.com/in/mykyta-kissa-684aa833a"

func (b *Basics) ToModel() cv.Basics {
	return cv.Basics{
		FullName:    b.FullName,
		Position:    b.Position,
		Description: b.Description,
		Email:       b.Email,
		Phone:       b.Phone,
		Links:       cv.Links(b.Links),
	}
}
