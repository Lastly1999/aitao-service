package sequencer

import "github.com/gin-gonic/gin"

type JsonResult struct {
	Msg     string       `json:"string"` // 响应消息
	Data    interface{}  `json:"data"`   // 响应数据
	Code    int          `json:"code"`   // 响应状态码
	Status  int          `json:"status"` // 自定义响应状态码
	Context *gin.Context // gin 上下文
}

// Send 返回json数据
func (jsonResult *JsonResult) Send(code, Status int, msg string, data interface{}) {
	jsonResult.Msg = msg
	jsonResult.Data = data
	jsonResult.Code = code
	jsonResult.Status = Status
	jsonResult.Context.JSON(code, gin.H{
		"msg":    jsonResult.Msg,
		"status": jsonResult.Status,
		"data":   jsonResult.Data,
	})
}
