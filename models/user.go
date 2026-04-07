package models

import "time"

type User struct {
	ID        uint           `gorm:"primary_key" json:"_id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"_"`
	FullName  string         `json:"full_name"`
	Email     string         `json:"email"`
}

func NewUser() *User {
	return &User{}
}
