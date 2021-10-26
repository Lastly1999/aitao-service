package test

import (
	"aitao-service/pkg/taobao_sdk"
	"aitao-service/service"
	"testing"
)

var commoditysService service.CommodityService

func Test_GetCommoditys(t *testing.T) {
	taobao_sdk.InitOpenTaobaoConf()
	commoditysService.GetCommoditys()
}
