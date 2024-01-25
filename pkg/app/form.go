package app

import (
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"go-gin-swagger/pkg/logging"
	"net/http"

	"go-gin-swagger/pkg/e"
)

// BindAndValid binds and validates data
func BindAndValid(c *gin.Context, form interface{}) (int, int) {
	err := c.Bind(form)
	if err != nil {
		return http.StatusBadRequest, e.INVALID_PARAMS
	}

	valid := validation.Validation{}
	check, err := valid.Valid(form)
	if err != nil {
		logging.Error("BindAndValid valid.Valid error", "msg", err.Error())
		return http.StatusInternalServerError, e.ERROR
	}
	if !check {
		MarkErrors(valid.Errors)
		return http.StatusBadRequest, e.INVALID_PARAMS
	}

	return http.StatusOK, e.SUCCESS
}
