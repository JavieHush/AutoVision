package dao

import (
	"AutoVision/internal/model"
	"AutoVision/pkg/app"
	"github.com/jinzhu/gorm"
	"time"
)

func (d *Dao) GetModelList(userID uint, page, pageSize int) ([]*model.Models, error) {
	// todo: some fields are nil here, will cause error?
	models := model.Models{UserID: userID}
	pageOffset := app.GetPageOffset(page, pageSize)
	return models.List(d.engine, pageOffset, pageSize)
}

func (d *Dao) GetModelByID(modelID uint) (*model.Models, error) {
	// todo: need error check here
	models := model.Models{Model: gorm.Model{ID: modelID}}

	return models.GetModel(d.engine)
}

func (d *Dao) CreateModel(
	modelPath string,
	modelSize float32,
	timeSpend int64,
	trainedAt time.Time,
	valAccuracy float32,
	dataSetID uint,
	userID uint,
) error {
	models := model.Models{
		Model:       gorm.Model{},
		ModelPath:   modelPath,
		ModelSize:   modelSize,
		TimeSpend:   timeSpend,
		TrainedAt:   trainedAt,
		ValAccuracy: valAccuracy,
		DatasetID:   dataSetID,
		UserID:      userID,
	}

	return models.Create(d.engine)
}

func (d *Dao) DeleteModel(modelID uint) error {
	models := model.Models{Model: gorm.Model{ID: modelID}}

	return models.Delete(d.engine)
}
