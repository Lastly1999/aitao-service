package routes

import (
	v1 "aitao-service/api/v1"
	"github.com/gin-gonic/gin"
)

func InitMaterialRoutes(router *gin.RouterGroup) {
	materialRoutes := router.Group("material")
	materialApi := v1.SysMaterialApi{}
	{
		materialRoutes.GET("material", materialApi.GetMaterialList)
	}
}
