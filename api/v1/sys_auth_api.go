package v1

import (
	"aitao-service/model/response"
	"aitao-service/service"
	"fmt"
	"github.com/gin-gonic/gin"
)

type AuthApi struct {
}

var authService service.AuthService

// WxLogin 授权登录
func (authApi *AuthApi) WxLogin(c *gin.Context) {
	checkParams := response.CheckAuthParams{}
	param := c.ShouldBindJSON(&checkParams)
	fmt.Println(param)
}
