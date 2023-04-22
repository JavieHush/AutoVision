package model

import (
	"gorm.io/gorm"
	"time"
)

type Datasets struct {
	gorm.Model
	Name        string    `json:"name"`
	Description string    `json:"description" gorm:"type:text"`
	DirPath     string    `json:"dir_path"`
	FileSize    int       `json:"file_size"`
	UploadTime  time.Time `json:"upload_time"`

	// 定义外键
	UserID uint  `json:"user_id"`
	Users  Users `gorm:"foreignKey:UserID"`
}

func (Datasets) TableName() string {
	return "datasets"
}
