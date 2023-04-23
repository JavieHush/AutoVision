package routers

import (
	v1 "AutoVision/internal/routers/api/v1"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	dataset := v1.NewDataset()
	model := v1.NewModel()
	apiV1 := r.Group("/api/v1")
	{
		apiV1.POST("/datasets", dataset.Upload)       // upload a dataset
		apiV1.DELETE("/datasets/:id", dataset.Delete) // delete a dataset by id
		apiV1.GET("/datasets", dataset.GetList)       // get list of dataset

		apiV1.DELETE("/models/:id", model.Delete)    // delete a model by id
		apiV1.GET("/models", model.GetList)          // get list of models
		apiV1.GET("/models/:id", model.GetModelByID) // get certain model by id
	}

	return r
}
