package input_json

import (
	"github.com/NikitaKissa/cvgen/internal/core"
	"github.com/NikitaKissa/cvgen/internal/cv"
)

type Project struct {
	Name        string         `json:"name"`
	Description *string        `json:"description"`
	Url         core.StringUrl `json:"url"`
}

func (p *Project) ToModel() cv.Project {
	return cv.Project{
		Name:        p.Name,
		Description: p.Description,
		Url:         *p.Url.Parse(),
	}
}

func projectsToModel(projects []Project) []cv.Project {
	var output = make([]cv.Project, len(projects))
	for i, item := range projects {
		output[i] = item.ToModel()
	}
	return output
}
