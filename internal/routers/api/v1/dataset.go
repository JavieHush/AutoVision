package v1

import (
	"AutoVision/global"
	"AutoVision/internal/service"
	"AutoVision/pkg/app"
	"AutoVision/pkg/convert"
	"AutoVision/pkg/errcode"
	"github.com/gin-gonic/gin"
)

type Dataset struct {
}

func NewDataset() Dataset {
	return Dataset{}
}

func (d Dataset) Delete(c *gin.Context) {
	param := service.DeleteDatasetRequest{DatasetID: convert.StrTo(c.Param("dataset_id")).MustUInt()}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	// if param is not valid.
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	err := svc.DeleteDataset(&param)

	if err != nil {
		global.Logger.Errorf("svc.DeleteDataset err: %v", err)
	}

	response.ToResponse(gin.H{})
}

func (d Dataset) GetTrainDatasetList(c *gin.Context) {
	param := service.DatasetTrainListRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	pager := app.Pager{Page: app.GetPage(c), PageSize: app.GetPageSize(c)}
	trainSet, err := svc.DatasetTrainList(&param, &pager)
	if err != nil {
		global.Logger.Errorf("svc.DatasetTrainList err: %v", err)
		response.ToErrorResponse(errcode.ErrorGetTrainListFail)
		return
	}

	response.ToResponseList(trainSet, 1)
}

func (d Dataset) GetPredDatasetList(c *gin.Context) {
	param := service.DatasetPredListRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	pager := app.Pager{Page: app.GetPage(c), PageSize: app.GetPageSize(c)}
	predSet, err := svc.DatasetPredList(&param, &pager)
	if err != nil {
		global.Logger.Errorf("svc.DatasetPredList err: %v", err)
		response.ToErrorResponse(errcode.ErrorGetPredListFail)
		return
	}

	response.ToResponseList(predSet, 1)
}

// UploadTrainedDataset todo: 这里还没有处理文件夹的本地保存
func (d Dataset) UploadTrainedDataset(c *gin.Context) {
	form, _ := c.MultipartForm()
	param := service.CreateTrainDatasetRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	err := svc.CreateTrainDataset(&param, len(form.File)-1)
	if err != nil {
		global.Logger.Errorf("svc.CreateTrainDataset err: %v", err)
		response.ToErrorResponse(errcode.ErrorCreateDatasetFail)
		return
	}

	response.ToResponse(gin.H{})
	return
}

func (d Dataset) UploadPredDataset(c *gin.Context) {
	form, _ := c.MultipartForm()
	param := service.CreatePredDatasetRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	err := svc.CreatePredDataset(&param, len(form.File)-1)
	if err != nil {
		global.Logger.Errorf("svc.CreateTrainDataset err: %v", err)
		response.ToErrorResponse(errcode.ErrorCreateDatasetFail)
		return
	}

	response.ToResponse(gin.H{})
	return
}
