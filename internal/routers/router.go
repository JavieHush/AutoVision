package routers

import (
	"AutoVision/internal/middleware"
	v1 "AutoVision/internal/routers/api/v1"
	"AutoVision/internal/service"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.InitCors())

	dataset := v1.NewDataset()
	modal := v1.NewModel()
	apiV1 := r.Group("/api/v1")
	{
		// 数据集
		apiV1.POST("/trainedDataset", dataset.UploadTrainedDataset)
		apiV1.POST("/predDataset", dataset.UploadPredDataset)

		apiV1.POST("/deleteDataset", dataset.Delete)

		apiV1.POST("/trainedDatasetList", dataset.GetTrainDatasetList)
		apiV1.POST("/predDatasetList", dataset.GetPredDatasetList)

		// 模型
		apiV1.POST("/modalList", modal.GetList)
		apiV1.POST("/createModal", modal.CreateModel)
		apiV1.POST("/deleteModal", modal.Delete)

		// 文件
		apiV1.GET("/resCSV", service.DownloadHandler)
	}

	return r
}
