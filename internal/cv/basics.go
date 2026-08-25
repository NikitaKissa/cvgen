package cv

import "net/url"

type Basics struct {
	Photo       *string
	FullName    string
	Position    string
	Description string
	Email       *string
	Phone       *string
	Links       Links
}

type Links map[string]url.URL // ex: "LinkedIn": "https://www.linkedin.com/in/mykyta-kissa-684aa833a"
