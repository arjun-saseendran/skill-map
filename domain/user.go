package main

import (
	"errors"

	"github.com/arjun-saseendran/skill-map/dto"
	"github.com/arjun-saseendran/skill-map/models"
)

type User interface {
	Create(userData dto.UserCreateInput) (*models.User, error)
	List() ([]models.User, error)
}
type user struct{}

func NewUser() User {
	return &user{}
}

func (u *user) Create(userData dto.UserCreateInput) (*models.User, error) {
	newUser := &models.User{FullName: userData.FullName, Email: userData.Email}
	db.DB.Create(newUser)
	if newUser.ID == 0 {
		return nil, errors.New("user creation failed")
	}
	return newUser, nil
}

func (u *user) List() ([]models.User, error) {
	users := []models.User{}
	db.DB.Find(&users)
	return users, nil
}

func (u *user) Get(string id) (*models.User, error) {
	singleUser := models.NewUser()
	db.DB

}
