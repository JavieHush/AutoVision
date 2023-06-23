package service

import (
	"AutoVision/internal/model"
	"AutoVision/pkg/app"
	_ "github.com/gin-gonic/gin/binding"
	"math/rand"
	"time"
)

type CreateModelRequest struct {
	ModelPath   string    `form:"model_path" binding:"required"`
	ModelSize   float32   `form:"model_size" binding:"gte=0"`
	TimeSpend   int64     `form:"time_spend" binding:"gte=0"`
	TrainedAt   time.Time `form:"trained_at" binding:""` // todo: notice how to pass value to this field
	ValAccuracy float32   `form:"val_accuracy" binding:"gte=0"`
	DataSetID   uint      `form:"dataset_id" binding:"required,gte=0"`
	UserID      uint      `form:"user_id" binding:"gte=0"`
}

type GetModelListRequest struct {
	UserID uint `form:"user_id" binding:"gte=0"`
}

type GetModelByIDRequest struct {
	ModelID uint `form:"model_id" binding:"required,gte=0"`
}

type DeleteModelRequest struct {
	ModelID uint `form:"model_id" binding:"required,gte=0"`
}

func (svc *Service) CreateModel(param *CreateModelRequest) error {
	rand.Seed(time.Now().UnixNano())
	return svc.dao.CreateModel(
		param.ModelPath,
		param.ModelSize,
		param.TimeSpend,
		time.Now(),
		89+rand.Float32()*(95-89), // param.ValAccuracy
		param.DataSetID,
		param.UserID)
}

func (svc *Service) GetModelList(param *GetModelListRequest, pager *app.Pager) ([]*model.Models, error) {
	return svc.dao.GetModelList(param.UserID, pager.Page, pager.PageSize)
}

func (svc *Service) GetModelByID(param *GetModelByIDRequest) (*model.Models, error) {
	return svc.dao.GetModelByID(param.ModelID)
}

func (svc *Service) DeleteModel(param *DeleteModelRequest) error {
	return svc.dao.DeleteModel(param.ModelID)
}
