package model

import "aitao-service/global"

type SysMaterials struct {
	global.GlobalModel
	MaterialId   string `json:"materialId" gorm:"material_id;comment:物料库id"`
	MaterialName string `json:"materialName" gorm:"Material_name;comment:物料库名称"`
}

func (SysMaterials) TableName() string {
	return "sys_materials"
}
