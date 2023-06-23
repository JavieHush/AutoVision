package dao

import (
	"AutoVision/internal/model"
	"AutoVision/pkg/app"
	"github.com/jinzhu/gorm"
	"time"
)

func (d *Dao) GetTrainDataset(userID uint, page, pagerSize int) ([]*model.Datasets, error) {
	trainDataset := model.Datasets{
		Model:       gorm.Model{},
		Name:        "",
		Description: "",
		DirPath:     "",
		FileSize:    0,
		UploadTime:  time.Time{},
		IsTrain:     1,
		UserID:      userID,
		Users:       model.Users{},
	}

	pageOffset := app.GetPageOffset(page, pagerSize)

	return trainDataset.DatasetTrainList(d.engine, pageOffset, pagerSize)
}

func (d *Dao) CreateTrainDataset(name string, description string, uploadTime time.Time, userID uint, fileSize int) error {
	trainDataset := model.Datasets{
		Model:       gorm.Model{},
		Name:        name,
		Description: description,
		//DirPath:     fmt.Sprintf("/%s", name),
		DirPath:    "",
		FileSize:   fileSize,
		UploadTime: uploadTime,
		IsTrain:    1,
		UserID:     userID,
		Users:      model.Users{},
	}

	return trainDataset.CreateDataset(d.engine)
}

func (d *Dao) DeleteDataset(id uint) error {
	dataset := model.Datasets{Model: gorm.Model{
		ID: id,
	}}

	return dataset.DeleteDataset(d.engine)
}

func (d *Dao) GetPredDataset(userID uint, page, pageSize int) ([]*model.Datasets, error) {
	predDataset := model.Datasets{
		Model:       gorm.Model{},
		Name:        "",
		Description: "",
		DirPath:     "",
		FileSize:    0,
		UploadTime:  time.Time{},
		IsTrain:     0,
		UserID:      userID,
		Users:       model.Users{},
	}

	pageOffset := app.GetPageOffset(page, pageSize)

	return predDataset.DatasetPredList(d.engine, pageOffset, pageSize)
}

func (d *Dao) CreatePredDataset(name string, description string, uploadTime time.Time, userID uint, fileSize int) error {
	trainDataset := model.Datasets{
		Model:       gorm.Model{},
		Name:        name,
		Description: description,
		//DirPath:     fmt.Sprintf("/%s", name),
		DirPath:    "",
		FileSize:   fileSize,
		UploadTime: uploadTime,
		IsTrain:    0,
		UserID:     userID,
		Users:      model.Users{},
	}

	return trainDataset.CreateDataset(d.engine)
}
