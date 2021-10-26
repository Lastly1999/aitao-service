package v1

import (
	"aitao-service/model/response"
	"aitao-service/model/sequencer"
	"aitao-service/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CommodityApi struct {
}

type ICommodityApi interface {
	GetCommoditys(c *gin.Context)
	GetRecommendCommoditys(c *gin.Context)
}

var commodityService service.CommodityService

// GetCommoditys TBK 商品物料搜索
func (commodityApi *CommodityApi) GetCommoditys(c *gin.Context) {
	jsonResult := sequencer.JsonResult{Context: c}
	// 参数绑定
	commoditParams := &response.CommoditParams{}
	err := c.ShouldBindJSON(commoditParams)
	if err != nil {
		jsonResult.Send(http.StatusOK, http.StatusBadGateway, err.Error(), nil)
		return
	}
	// 业务处理
	res, err := commodityService.GetCommoditys(commoditParams)
	if err != nil {
		jsonResult.Send(http.StatusOK, http.StatusOK, err.Error(), nil)
		return
	}
	jsonResult.Send(http.StatusOK, http.StatusOK, "ok", gin.H{
		"commodits": res.Get("tbk_dg_material_optional_response").Get("result_list").Get("map_data"),
		"total":     res.Get("tbk_dg_material_optional_response").Get("total_results"),
	})
}

// GetRecommendCommoditys TBK 精选商品物料搜索
func (commodityApi *CommodityApi) GetRecommendCommoditys(c *gin.Context) {
	jsonResult := sequencer.JsonResult{Context: c}
	res, err := commodityService.GetRecommendCommoditys()
	if err != nil {
		jsonResult.Send(http.StatusOK, http.StatusOK, err.Error(), nil)
		return
	}
	jsonResult.Send(http.StatusOK, http.StatusOK, "ok", gin.H{
		"data": res.Get("tbk_dg_optimus_material_response").Get("result_list").Get("map_data"),
	})
}
