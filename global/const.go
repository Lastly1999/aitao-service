package global

import (
	"github.com/silenceper/wechat/v2/miniprogram"
	"gorm.io/gorm"
)

var (
	GLOBAL_DB          *gorm.DB
	GLOBAL_MINI_CLIENT *miniprogram.MiniProgram
)
