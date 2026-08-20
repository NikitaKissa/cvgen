package input_json

type CV struct {
	Basics       Basics
	Education    []Education
	Experience   []Experience
	SkillGroups  []SkillGroup
	Certificates []Certificate
	Projects     []Project
	Languages    []Language
}
