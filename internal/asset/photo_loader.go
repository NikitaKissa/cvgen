package asset

import (
	"fmt"
	"net/url"
)

func LoadPhoto(input string) (string, error) {
	if input == "" {
		return "", nil
	}
	if isRemoteURL(input) {
		return input, nil
	}
	a, err := Load(input)
	if err != nil {
		return "", fmt.Errorf("load photo: %w", err)
	}
	return a.DataURI(), nil
}

func isRemoteURL(input string) bool {
	u, err := url.Parse(input)
	if err != nil {
		return false
	}
	return u.Scheme == "http" || u.Scheme == "https"
}
