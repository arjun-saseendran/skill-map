package domain

import (
	"errors"

	"github.com/arjun-saseendran/skill-map/db"
	"github.com/arjun-saseendran/skill-map/dto"
	"github.com/arjun-saseendran/skill-map/models"
)

type User interface {
	Create(userData dto.UserCreateInput) (*models.User, error)
	List() ([]models.User, error)
	Get(id string) (*models.User, error)
	Update(id string, userData *dto.UserUpdateInput) (*models.User, error)
	Delete(id string) error
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
	var users []models.User
	db.DB.Find(&users)
	return users, nil
}

func (u *user) Get(id string) (*models.User, error) {
	singleUser := models.NewUser()
	db.DB.First(&singleUser, id)
	if singleUser.ID == 0 {
		return nil, errors.New("no user found")
	}
	return singleUser, nil
}

func (u *user) Update(id string, userData *dto.UserUpdateInput) (*models.User, error) {
	updateUser := models.NewUser()
	db.DB.First(updateUser, id)
	if updateUser.ID == 0 {
		return nil, errors.New("user not found")
	}
	db.DB.Model(&updateUser).Updates(models.User{FullName: userData.FullName, Email: userData.Email})
	return updateUser, nil
}

func (u *user) Delete(id string) error {
	deleteUser := models.NewUser()
	db.DB.First(deleteUser, id)
	if deleteUser.ID == 0 {
		return errors.New("user not found")
	}
	db.DB.Delete(deleteUser)
	return nil
}
