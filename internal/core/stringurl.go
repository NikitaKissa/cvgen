package core

import "net/url"

type StringUrl string

func (u *StringUrl) Parse() *url.URL {
	out, err := url.Parse(string(*u))
	if err != nil {
		return &url.URL{
			Scheme: "http",
			Host:   "invalid.url",
			RawQuery: url.Values{
				"err": []string{err.Error()},
			}.Encode(),
		}
	}
	return out
}
