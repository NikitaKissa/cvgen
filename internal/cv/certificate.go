package cv

import "net/url"

type Certificate struct {
	Name   string
	Issuer string
	Url    url.URL
}
