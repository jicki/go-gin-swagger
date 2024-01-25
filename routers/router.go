package routers

import (
	"github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "go-gin-swagger/docs"
	"go-gin-swagger/pkg/logging"
	"net/http"
	"time"
)

var r *gin.Engine

func InitRouter() *gin.Engine {

	r = gin.New()
	r.Use(ginzap.Ginzap(logging.Logger, time.RFC3339, true))
	r.Use(ginzap.RecoveryWithZap(logging.Logger, true))

	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Welcome to gin-swagger-demo")
	})

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "pong",
		})
	})

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	InitAppRouter()
	InitAdminRouter()

	return r
}
