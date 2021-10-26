package routes

import (
	v1 "aitao-service/api/v1"
	"github.com/gin-gonic/gin"
)

func InitCommodityRoutes(router *gin.RouterGroup) {
	commodityRoutes := router.Group("commodity")
	commodityApi := v1.CommodityApi{}
	{
		// 获取物料商品
		commodityRoutes.POST("commodity", commodityApi.GetCommoditys)
		// 获取精选物料商品
		commodityRoutes.POST("featured", commodityApi.GetRecommendCommoditys)
	}
}
