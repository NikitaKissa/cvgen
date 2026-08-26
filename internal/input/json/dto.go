package input_json

import (
	"github.com/NikitaKissa/cvgen/internal/cv"
)

type InputDTO struct {
	Basics       Basics        `json:"basics"`
	Education    []Education   `json:"education"`
	Experience   []Experience  `json:"experience"`
	SkillGroups  []SkillGroup  `json:"skill_groups"`
	Certificates []Certificate `json:"certificates"`
	Projects     []Project     `json:"projects"`
	Languages    []Language    `json:"languages"`
	Clause       *string       `json:"clause"`
}

// for future
func (dto *InputDTO) Validate() error {
	return nil
}

func (dto *InputDTO) ToModel() cv.CV {
	return cv.CV{
		Basics:       dto.Basics.ToModel(),
		Education:    educationToModel(dto.Education),
		Experience:   experienceToModel(dto.Experience),
		SkillGroups:  skillGroupsToModel(dto.SkillGroups),
		Certificates: certificatesToModel(dto.Certificates),
		Projects:     projectsToModel(dto.Projects),
		Languages:    languagesToModel(dto.Languages),
		Clause:       dto.Clause,
	}
}
