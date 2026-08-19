package cv

import (
	"net/url"
	"time"
)

type Education struct {
	Institution        string
	InstitutionWebsite *url.URL
	FieldOfStudy       string
	Description        *string
	From               time.Time
	To                 *time.Time // nil means "till now"
}
