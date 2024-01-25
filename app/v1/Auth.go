package v1

import (
	"github.com/gin-gonic/gin"
	"go-gin-swagger/pkg/logging"
	"net/http"

	"go-gin-swagger/pkg/app"
	"go-gin-swagger/pkg/e"
	"go-gin-swagger/pkg/util"
)

func RefreshToken(c *gin.Context) {
	appG := app.Gin{C: c}
	userId := appG.GetUid()

	token, err := util.GenerateUserToken(userId)
	if err != nil {
		logging.Error("RefreshToken err:", "msg", err.Error(), "userId", userId)
		appG.Response(http.StatusInternalServerError, e.ERROR_AUTH_TOKEN, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, token)
}
