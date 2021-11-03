package service

import (
	"aitao-service/global"
	"github.com/silenceper/wechat/v2/miniprogram/auth"
)

type AuthService struct {
}
type IAuthService interface {
	GetUserAuth(jsCode string) (result auth.ResCode2Session, err error)
}

// GetUserAuth 获取用户信息
func (authService *AuthService) GetUserAuth(jsCode string) (code2Session auth.ResCode2Session, err error) {
	auth := global.GLOBAL_MINI_CLIENT.GetAuth()
	code2Session, err = auth.Code2Session(jsCode)
	if err != nil {
		return code2Session, err
	}
	return code2Session, nil
}
