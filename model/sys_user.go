package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserName string `json:"userName" gorm:"column:user_name"`
	PassWord string `json:"passWord" gorm:"column:pass_word"`
	OpenId   string `json:"openId" gorm:"column:open_id"`
}
