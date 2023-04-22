package model

import "gorm.io/gorm"

type Users struct {
	gorm.Model
	Username string `json:"username" gorm:"uniqueIndex"`
	Password string `json:"password"`
	Email    string `json:"email" gorm:"uniqueIndex"`
}

func (Users) TableName() string {
	return "users"
}
