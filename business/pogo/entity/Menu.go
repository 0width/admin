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
	Type     int    `gorm:"type:tinyint;comment:1:目录 2:菜单 3:按钮"`
	Status   int    `gorm:"type:tinyint;default:0;not null;commnet:0: 正常 1: 停用"`
	Visible  int    `gorm:"type:tinyint;default:0;not null;comment:0: 正常 1: 隐藏"`
	Query    string `gorm:"type:string;size:1023;comment:路由参数"`
}

func (Menu) TableName() string {
	return "menu"
}
