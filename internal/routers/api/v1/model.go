package v1

// 新增或者删除该文件中的函数后，要同步修改 internal/routers下的router.go文件
// todo: current status: finish this file.

import (
	"AutoVision/global"
	"AutoVision/internal/service"
	"AutoVision/pkg/app"
	"AutoVision/pkg/convert"
	"AutoVision/pkg/errcode"
	"github.com/gin-gonic/gin"
)

type Model struct {
}

func NewModel() Model { // create a model object
	return Model{}
}

func (m Model) GetList(c *gin.Context) {
	param := service.GetModelListRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	// param is not valid.
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	pager := app.Pager{Page: app.GetPage(c), PageSize: app.GetPageSize(c)}
	// here is different from tutorial.(2.6.7)

	models, err := svc.GetModelList(&param, &pager)
	// if error occurred
	if err != nil {
		global.Logger.Errorf("svc.GetModelList err: %v", err)
		response.ToErrorResponse(errcode.ErrorGetModelList)
		return
	}

	// totalRows is not set.
	response.ToResponseList(models, 0)
	return
}

func (m Model) Delete(c *gin.Context) {
	param := service.DeleteModelRequest{ModelID: convert.StrTo(c.Param("model_id")).MustUInt()}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	// if param is not valid.
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	err := svc.DeleteModel(&param)
	// if error occurred.
	if err != nil {
		global.Logger.Errorf("svc.DeleteModel err: %v", err)
	}

	response.ToResponse(gin.H{})

}

func (m Model) GetModelByID(c *gin.Context) {
	param := service.GetModelByIDRequest{ModelID: convert.StrTo(c.Param("model_id")).MustUInt()}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	// if param is not valid.
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	model, err := svc.GetModelByID(&param)
	// if error occurred.
	if err != nil {
		global.Logger.Errorf("svc.GetModelByID err: %v", err)
	}

	response.ToResponse(model)
}

func (m Model) CreateModel(c *gin.Context) {
	param := service.CreateModelRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	err := svc.CreateModel(&param)
	if err != nil {
		global.Logger.Errorf("svc.CreateModel err: %v", err)
	}

	response.ToResponse(gin.H{})
}
