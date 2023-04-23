package v1

import "github.com/gin-gonic/gin"

type Dataset struct {
}

func NewDataset() Dataset {
	return Dataset{}
}

func (d Dataset) GetList(c *gin.Context) {

}

func (d Dataset) Upload(c *gin.Context) {

}

func (d Dataset) Delete(c *gin.Context) {

}
