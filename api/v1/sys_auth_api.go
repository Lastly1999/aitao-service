package v1

import (
	"aitao-service/model/response"
	"aitao-service/model/sequencer"
	"aitao-service/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthApi struct {
}

var authService service.AuthService

// WxLogin 授权登录
func (authApi *AuthApi) WxLogin(c *gin.Context) {
	jsonResult := sequencer.JsonResult{Context: c}
	checkParams := response.CheckAuthParams{}
	err := c.ShouldBindJSON(&checkParams)
	if err != nil {
		jsonResult.Send(http.StatusOK, http.StatusInternalServerError, "json解析失败", nil)
		return
	}
	auth, err := authService.GetUserAuth(checkParams.JsCode)
	if err != nil {
		jsonResult.Send(http.StatusOK, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	jsonResult.Send(http.StatusOK, http.StatusOK, "请求成功", auth)
}
