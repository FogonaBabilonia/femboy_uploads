package models

import (
	"time"

	"github.com/FogonaBabilonia/femboy_uploads/services"
	"gorm.io/gorm"
)

type base_user struct {
	Name     string `form:"name" binding:"required,min=4,max=36"`
	Password string `form:"password" binding:"required,min=8,max=36"`
}

type model struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

type User struct {
	model
	Name     string `json:"name"`
	Password string `json:"-"`
}

type Create_User struct {
	base_user
}

type SignIn_User struct {
	base_user
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
