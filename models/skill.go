package models

import (
	"time"

	"gorm.io/gorm"
)

type Skill struct {
	ID        uint           `gorm:"primary_key" json:"_id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"_"`
	Name      string         `json:"name"`
}

type SkillGroup struct {
	ID        uint           `gorm:"primary_key" json:"_id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"_"`
	Name      string         `json:"name"`
	Skills    []Skill        `gorm:"many2many:skillgroup_skills;" json:"skills"`
}

func NewSkill() *Skill {
	return &Skill{}
}

func NewSkillGroup() *SkillGroup {
	return &SkillGroup{}
}
