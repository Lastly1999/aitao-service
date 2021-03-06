package main

import (
	"aitao-service/pkg/config"
	"aitao-service/pkg/gorm"
	"aitao-service/pkg/taobao_sdk"
	"aitao-service/pkg/wechat_sdk"
	"aitao-service/router"
)

func init() {
	// sys config
	config.InitServerConfig()
	// gorm
	gorm.InitGormDb()
	// 微信小程序sdk
	wechat_sdk.InitWeClient()
	// 淘宝sdk
	taobao_sdk.InitOpenTaobaoConf()
}

func main() {
	err := router.InitRouter().Run(":" + config.GLOBAL_CONF.Server.Port)
	if err != nil {
		return
	}
}
