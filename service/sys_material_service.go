package service

import (
	"aitao-service/global"
	"aitao-service/model"
)

type MaterialService struct {
}

type IMaterialService interface {
	GetMaterials() (materials []*model.SysMaterials, err error)
}

// GetMaterials 获取物料库分类
func (materialService *MaterialService) GetMaterials() (materials []*model.SysMaterials, err error) {
	if err = global.GLOBAL_DB.Find(&materials).Error; err != nil {
		return nil, err
	}
	return materials, nil
}
