package model

import (
	"aitao-service/global"
)

type SysUser struct {
	global.GlobalModel
	UserName string `json:"userName" gorm:"column:user_name"`
	PassWord string `json:"passWord" gorm:"column:pass_word"`
	OpenId   string `json:"openId" gorm:"column:open_id"`
}

func (SysUser) TableName() string {
	return "sys_users"
}
