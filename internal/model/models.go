package model

import (
	"gorm.io/gorm"
	"time"
)

type Models struct {
	gorm.Model
	ModelPath   string    `json:"model_path"`
	ModelSize   float32   `json:"model_size"`
	TimeSpend   int64     `json:"time_spend"`
	TrainedAt   time.Time `json:"trained_at"`
	ValAccuracy float32   `json:"val_accuracy"`

	// 外键
	DatasetID uint     `json:"dataset_id"`
	UserID    uint     `json:"user_id"`
	Datasets  Datasets `gorm:"foreignKey:DatasetID"`
	Users     Users    `gorm:"foreignKey:UserID"`
}

func (Models) TableName() string {
	return "models"
}
