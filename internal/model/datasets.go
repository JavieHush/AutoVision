package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Datasets struct {
	gorm.Model
	Name        string    `json:"name"`
	Description string    `json:"description" gorm:"type:text"`
	DirPath     string    `json:"dir_path"`
	FileSize    int       `json:"file_size"`
	UploadTime  time.Time `json:"upload_time"`
	IsTrain     int       `json:"is_train"`

	// 定义外键
	UserID uint  `json:"user_id"`
	Users  Users `gorm:"foreignKey:UserID"`
}

func (Datasets) TableName() string {
	return "datasets"
}

/*
	func for users.
*/

func (d Datasets) DeleteDataset(db *gorm.DB) error {
	return db.Where("deleted_at is null").Where("ID = ? ", d.Model.ID).Delete(&d).Error // DeletedAt field will be automatically set to current time
}

// DatasetTrainList DatasetList List get dataset list
func (d Datasets) DatasetTrainList(db *gorm.DB, pageOffset, pageSize int) ([]*Datasets, error) {
	var datasets []*Datasets
	var err error
	if pageOffset >= 0 && pageSize > 0 {
		db = db.Offset(pageOffset).Limit(pageSize)
	}
	// find certain user's all models that are not deleted by userID.
	// todo: is it right here using UserID?
	if err = db.Where("deleted_at is null").Where("user_id = ? and is_train = ?", d.UserID, 1).Preload("Users").Find(&datasets).Error; err != nil {
		return nil, err
	}

	return datasets, nil
}

func (d Datasets) DatasetPredList(db *gorm.DB, pageOffset, pageSize int) ([]*Datasets, error) {
	var datasets []*Datasets
	var err error
	if pageOffset >= 0 && pageSize > 0 {
		db = db.Offset(pageOffset).Limit(pageSize)
	}
	// find certain user's all models that are not deleted by userID.
	// todo: is it right here using UserID?
	if err = db.Where("deleted_at is null").Where("user_id = ? and is_train = ?", d.UserID, 0).Preload("Users").Find(&datasets).Error; err != nil {
		return nil, err
	}

	return datasets, nil
}

// GetDataset GetModel get certain model by ID
// todo: notice return value should be *Datasets or Datasets
func (d Datasets) GetDataset(db *gorm.DB) (*Datasets, error) {
	var target Datasets
	var err error
	if err = db.Where("deleted_at is null").Where("ID = ?", d.ID).First(&target).Error; err != nil {
		return &Datasets{}, err
	}
	return &target, nil
}

/*
	func for sys
*/

func (d Datasets) CreateDataset(db *gorm.DB) error {
	return db.Create(&d).Error
}
