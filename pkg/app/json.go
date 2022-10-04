package app

import (
	"github.com/astaxie/beego/validation"
	"github.com/coolishbee/go-gin-sample/pkg/e"
	"github.com/gin-gonic/gin"
)

// BindAndValid binds and validates data
func BindAndValid(c *gin.Context, json interface{}) int {
	err := c.Bind(json)
	if err != nil {
		return e.INVALID_PARAMS
	}

	valid := validation.Validation{}
	check, err := valid.Valid(json)
	if err != nil {
		return e.ERROR
	}
	if !check {
		MarkErrors(valid.Errors)
		return e.INVALID_PARAMS
	}

	return e.SUCCESS
}
