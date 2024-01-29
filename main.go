package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go-gin-swagger/models"
	"go-gin-swagger/pkg/gredis"
	"go-gin-swagger/pkg/logging"
	"go-gin-swagger/pkg/setting"
	"go-gin-swagger/pkg/util"
	"go-gin-swagger/routers"
)

func init() {
	setting.Setup()
	logging.Setup()
	gredis.Setup()
	util.Setup()
	models.Setup()
}

// @title go-gin-swagger
// @version 1.0
// @description go-gin-swagger 服务后端API接口文档
// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @host localhost:8000
// @BasePath
func main() {
	gin.SetMode(setting.ServerSetting.RunMode)

	routersInit := routers.InitRouter()
	readTimeout := setting.ServerSetting.ReadTimeout
	writeTimeout := setting.ServerSetting.WriteTimeout
	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)
	maxHeaderBytes := 1 << 20

	server := &http.Server{
		Addr:           endPoint,
		Handler:        routersInit,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}

	log.Printf("[info] start http server listening %s", endPoint)

	server.ListenAndServe()
}
