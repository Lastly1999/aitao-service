package routes

import (
	v1 "aitao-service/api/v1"
	"github.com/gin-gonic/gin"
)

func InitAuthRoutes(router *gin.RouterGroup) {
	authRoutes := router.Group("auth")
	authApi := v1.AuthApi{}
	{
		// 用户登录
		authRoutes.POST("login", authApi.WxLogin)
	}
}
