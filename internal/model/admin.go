package model

import "gorm.io/gorm"

type Admin struct {
	gorm.Model
	Name     string `json:"name" gorm:"uniqueIndex"`
	Password string `json:"password"`
}

func (Admin) TableName() string {
	return "admin"
}
