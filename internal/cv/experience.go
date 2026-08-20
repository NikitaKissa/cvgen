package cv

import (
	"net/url"
	"time"
)

type Experience struct {
	Company        string
	CompanyWebsite *url.URL
	Position       string
	Description    *string
	Stack          []string
	From           time.Time
	To             *time.Time // nil means "till now"
}
