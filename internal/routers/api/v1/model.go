package v1

import "github.com/gin-gonic/gin"

type Model struct {
}

func NewModel() Model { // create a model object
	return Model{}
}

func (m Model) GetList(c *gin.Context) {

}

func (m Model) Delete(c *gin.Context) {

}

func (m Model) GetModelByID(c *gin.Context) {

}
