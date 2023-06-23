package model

import (
	"github.com/jinzhu/gorm"
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

func (Models) TableName() string { // TableName is automatically used by gorm
	return "models"
}

/*
	func for users.
*/

func (m Models) Delete(db *gorm.DB) error {
	return db.Where("deleted_at is null").Where("id = ? ", m.Model.ID).Delete(&m).Error // DeletedAt field will be automatically set to current time
}

// List get model list
func (m Models) List(db *gorm.DB, pageOffset, pageSize int) ([]*Models, error) {
	var models []*Models
	var err error
	if pageOffset >= 0 && pageSize > 0 {
		db = db.Offset(pageOffset).Limit(pageSize)
	}
	// find certain user's all models that are not deleted by userID.
	// todo: is it right here using UserID?
	if err = db.Where("deleted_at is null").Where("user_id = ?", m.UserID).Preload("Users").Find(&models).Error; err != nil {
		return nil, err
	}

	return models, nil
}

// GetModel get certain model by ID
// todo: notice return value should be *Models or Models
func (m Models) GetModel(db *gorm.DB) (*Models, error) {
	var target Models
	var err error
	if err = db.Where("deleted_at is null").Where("id = ?", m.ID).First(&target).Error; err != nil {
		return &Models{}, err
	}
	return &target, nil
}

/*
	func for sys
*/

func (m Models) Create(db *gorm.DB) error {
	return db.Create(&m).Error
}
