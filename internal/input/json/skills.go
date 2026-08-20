package input_json

import "github.com/NikitaKissa/cvgen/internal/cv"

type SkillGroup struct {
	Name   string   `json:"name"`
	Skills []string `json:"skills"`
}

func (sg *SkillGroup) ToModel() cv.SkillGroup {
	return cv.SkillGroup{
		Name:   sg.Name,
		Skills: sg.Skills,
	}
}

func skillGroupsToModel(skillGroup []SkillGroup) []cv.SkillGroup {
	var output = make([]cv.SkillGroup, len(skillGroup))
	for i, item := range skillGroup {
		output[i] = item.ToModel()
	}
	return output
}
