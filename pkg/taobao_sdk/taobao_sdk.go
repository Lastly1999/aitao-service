package taobao_sdk

import (
	"aitao-service/pkg/config"
	"github.com/nilorg/go-opentaobao"
)

func InitOpenTaobaoConf() {
	opentaobao.AppKey = config.GLOBAL_CONF.TaoSdk.AppKey
	opentaobao.AppSecret = config.GLOBAL_CONF.TaoSdk.AppSecret
	opentaobao.Router = config.GLOBAL_CONF.TaoSdk.Router
}
