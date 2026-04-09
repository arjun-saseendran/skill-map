package models

import (
	"time"

	"gorm.io/gorm"
)

type UserSkillRank struct {
	ID        uint           `gorm:"primarykey" json:"_id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"_"`
	UserID    int
	User      User
	SkillID   int
	Skill     Skill
	Rank      uint `json:"rank"`
}

func NewUserSkillRank() *UserSkillRank {
	return &UserSkillRank{}
}
