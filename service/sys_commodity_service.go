package service

import (
	"aitao-service/model/request"
	"aitao-service/pkg/config"
	"github.com/bitly/go-simplejson"
	"github.com/nilorg/go-opentaobao"
)

type CommodityService struct {
}

type ICommodityService interface {
	GetCommoditys(commoditParams *request.CommoditParams) (res *simplejson.Json, err error)
	GetRecommendCommoditys(recommendParams *request.RecommendParams) (res *simplejson.Json, err error)
}

// GetCommoditys 物料商品搜索
func (commodityService *CommodityService) GetCommoditys(commoditParams *request.CommoditParams) (res *simplejson.Json, err error) {
	res, err = opentaobao.Execute("taobao.tbk.dg.material.optional", opentaobao.Parameter{
		"adzone_id": config.GLOBAL_CONF.TaoSdk.Adzoneid,
		"q":         commoditParams.Keywords,
		"page_size": commoditParams.PageSize,
		"page_no":   commoditParams.PageOn,
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}

// GetRecommendCommoditys 精选物料商品推荐查询
func (commodityService *CommodityService) GetRecommendCommoditys(recommendParams *request.RecommendParams) (res *simplejson.Json, err error) {
	res, err = opentaobao.Execute("taobao.tbk.dg.optimus.material", opentaobao.Parameter{
		"adzone_id":   config.GLOBAL_CONF.TaoSdk.Adzoneid,
		"page_size":   recommendParams.PageSize,
		"page_no":     recommendParams.PageOn,
		"material_id": recommendParams.MaterialId,
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}
