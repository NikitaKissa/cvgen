package asset

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
)

type Asset struct {
	Name        string
	ContentType string
	Data        []byte
}

func Load(path string) (Asset, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return Asset{}, fmt.Errorf("read asset file: %w", err)
	}
	return Asset{
		Name:        filepath.Base(path),
		ContentType: http.DetectContentType(data),
		Data:        data,
	}, nil
}

func (a Asset) DataURI() string {
	encoded := base64.StdEncoding.EncodeToString(a.Data)
	return fmt.Sprintf("data:%s;base64,%s", a.ContentType, encoded)
}
