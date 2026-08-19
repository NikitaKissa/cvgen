package cv

import "net/url"

type Project struct {
	Name        string
	Description *string
	Url         url.URL
}
