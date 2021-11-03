package global

import (
	"github.com/silenceper/wechat/v2/miniprogram"
	"gorm.io/gorm"
	"time"
)

var (
	GLOBAL_DB          *gorm.DB
	GLOBAL_MINI_CLIENT *miniprogram.MiniProgram
)

type GlobalModel struct {
	ID        uint           `json:"id" gorm:"primarykey"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}
