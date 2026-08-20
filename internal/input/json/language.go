package input_json

import "github.com/NikitaKissa/cvgen/internal/cv"

type Language struct {
	Name        string `json:"name"`
	Proficiency string `json:"proficiency"`
}

func (l *Language) ToModel() cv.Language {
	return cv.Language{
		Name:        l.Name,
		Proficiency: l.Proficiency,
	}
}

func languagesToModel(languages []Language) []cv.Language {
	var output = make([]cv.Language, len(languages))
	for i, item := range languages {
		output[i] = item.ToModel()
	}
	return output
}
