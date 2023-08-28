package models

import "gorm.io/gorm"

type User struct {
	*gorm.Model
	Name     string
	Password string
}

type Create_User struct {
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
}
