package taobao_sdk

import (
	"aitao-service/pkg/config"
	"fmt"
	"github.com/nilorg/go-opentaobao"
)

func InitOpenTaobaoConf() {
	fmt.Println(config.GLOBAL_CONF.TaoSdk)
	opentaobao.AppKey = config.GLOBAL_CONF.TaoSdk.AppKey
	opentaobao.AppSecret = config.GLOBAL_CONF.TaoSdk.AppSecret
	opentaobao.Router = "http://gw.api.taobao.com/router/rest"
}
