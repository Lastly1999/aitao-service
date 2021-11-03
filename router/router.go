package router

import (
	"aitao-service/middleware/cors"
	"aitao-service/router/routes"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	gin.ForceConsoleColor()
	app := gin.Default()
	// 跨域中间件
	app.Use(cors.Cors())
	// v1 apis
	apiRouter := app.Group("v1")
	{
		// 授权模块
		routes.InitAuthRoutes(apiRouter)
		// 商品模块
		routes.InitCommodityRoutes(apiRouter)
		// 物料库分类模块
		routes.InitMaterialRoutes(apiRouter)
	}
	return app
}
