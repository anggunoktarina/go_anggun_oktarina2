package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID           uint           `json:"id" form:"id" gorm:"primary_key"`
	Name         string         `json:"name" form:"name"`
	Email        string         `json:"email" form:"email"`
	Password     string         `json:"password" form:"password"`
	Token        string         `gorm:"-"`
}