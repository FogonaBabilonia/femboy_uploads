package models

import (
	"github.com/FogonaBabilonia/femboy_uploads/services"
	"gorm.io/gorm"
)

type User struct {
	*gorm.Model
	Name     string `json:"name"`
	Password string `json:"-"`
}

type Create_User struct {
	Name     string `form:"name" binding:"required"`
	Password string `form:"password" binding:"required"`
}

func NewUser(name, password string) (*User, error) {
	hashPass, err := services.GeneratePassword(password)

	if err != nil {
		return nil, err
	}

	return &User{
		Name:     name,
		Password: hashPass,
	}, nil
}
