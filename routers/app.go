package routers

import (
	v1 "go-gin-swagger/app/v1"
	"go-gin-swagger/middleware/jwt"
)

/*
	InitAppRouter

初始化 APP 接口的路径
*/
func InitAppRouter() {
	appv1 := r.Group("/app/v1")

	appv1.Use(jwt.JWT())
	{
		//refresh token
		appv1.GET("/auth/refreshToken", v1.RefreshToken)
	}
}
