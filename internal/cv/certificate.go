package cv

import "net/url"

type Certificate struct {
	Name   string
	Issuer string
	url    url.URL
}
