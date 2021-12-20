package systemEntity

import "gorm.io/gorm"

// SystemDeptEntity 部门表
type SystemDeptEntity struct {
	gorm.Model
	ParentId  *uint
	Parent    *SystemDeptEntity
	Ancestors string `gorm:"size:64;comment:祖级列表"`
	Name      string `gorm:"size:64"`
	Order     int    `gorm:"size:32"`
	Leader    string `gorm:"size:64"`
	Phone     string `gorm:"size:16"`
	Status    int    `gorm:"type:tinyint"`
}

func (SystemDeptEntity) TableName() string {
	return "dept"
}
