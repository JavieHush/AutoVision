package app

import (
	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin/binding"
	val "github.com/go-playground/validator/v10"
	"strings"
)

type ValidError struct {
	Key     string
	Message string
}

type ValidErrors []*ValidError

func (v *ValidError) Error() string {
	return v.Message
}

func (v ValidErrors) Error() string {
	return strings.Join(v.Errors(), ",")
}

func (v ValidErrors) Errors() []string {
	var errs []string
	for _, err := range v {
		errs = append(errs, err.Error())
	}

	return errs
}

// BindAndValid different from tutorial(in 2.5)
func BindAndValid(c *gin.Context, v interface{}) (bool, ValidErrors) {
	// todo: 接受到的请求为空
	var errs ValidErrors
	err := c.ShouldBind(v)
	if err != nil {
		_, ok := err.(val.ValidationErrors)
		if !ok {
			return false, errs
		}

		return false, errs
	}

	return true, nil
}
