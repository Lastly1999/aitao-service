package wechat_sdk

import (
	"aitao-service/global"
	"aitao-service/pkg/config"
	"github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/cache"
	miniConf "github.com/silenceper/wechat/v2/miniprogram/config"
)

// InitWeClient 初始化微信小程序sdk
func InitWeClient() {
	wc := wechat.NewWechat()
	memory := cache.NewMemory()
	conf := &miniConf.Config{
		AppID:     config.GLOBAL_CONF.WxSdk.Appid,
		AppSecret: config.GLOBAL_CONF.WxSdk.AppSecret,
		Cache:     memory,
	}
	global.GLOBAL_MINI_CLIENT = wc.GetMiniProgram(conf)
}
