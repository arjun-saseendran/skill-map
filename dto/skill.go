package dto

type SkillCreateInput struct {
	Name string `json:"name"`
}

type SkillUpdateInput struct {
	Name string `json:"name"`
}

type SkillGroupCreateInput struct {
	Name string `json:"name"`
}

type SkillGroupUpdateInput struct {
	Name   string `json:"name"`
	Skills []int  `json:"skills"`
}

func NewSkillCreateInput() *SkillCreateInput {
	return &SkillCreateInput{}
}

func NewSkillUpdateInput() *SkillUpdateInput {
	return &SkillUpdateInput{}
}

func NewSkillGroupCreateInput() *SkillGroupCreateInput {
	return &SkillGroupCreateInput{}
}
func NewSkillGroupUpdateInput() *SkillGroupUpdateInput {
	return &SkillGroupUpdateInput{}
}
