package entity

import (
	"gorm.io/gorm"
)

type Menu struct {
	gorm.Model
	Name     string `gorm:"type:string;size:64;comment:菜单名称"`
	ParentId string `gorm:"type:uint;comment:父菜单ID"`
	Remark   string `gorm:"type:string;size:255;comment:备注"`
	Perms    string `gorm:"type:string;size:64;comment:权限标识"`
	Type     string `gorm:"type:tinyint;comment:1:目录 2:菜单 3:按钮"`
}

func (Menu) TableName() string {
	return "menu"
}
