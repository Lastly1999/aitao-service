package v1

import (
	"aitao-service/model/sequencer"
	"aitao-service/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type SysMaterialApi struct {
}

type ISysMaterialApi interface {
	GetMaterialList(c *gin.Context)
}

var materialService service.MaterialService

func (sysMaterialApi *SysMaterialApi) GetMaterialList(c *gin.Context) {
	jsonResult := sequencer.JsonResult{Context: c}
	materials, err := materialService.GetMaterials()
	if err != nil {
		jsonResult.Send(http.StatusOK, http.StatusBadGateway, err.Error(), nil)
		return
	}
	jsonResult.Send(http.StatusOK, http.StatusOK, "ok", materials)
}
