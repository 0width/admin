package systemEntity

import (
	"gorm.io/gorm"
)

type Menu struct {
	gorm.Model
	Name      string `gorm:"type:string;size:64;comment:菜单路由名称"`
	Title     string `gorm:"type:string;size:64;comment:菜单名称"`
	Icon      string `gorm:"type:string;size:64;comment:图标"`
	Path      string `gorm:"type:string;size:64;comment:路由path"`
	Query     string `gorm:"type:string;size:511;comment:路由参数"`
	Redirect  string `gorm:"type:string;size:64;comment:路由redirect"`
	Component string `gorm:"type:string;size:64;comment:组件"`
	Order     int    `gorm:"type:int;size:32;comment:菜单顺序"`
	ParentId  uint   `gorm:"comment:父菜单ID"`
	Parent    *Menu
	Remark    string  `gorm:"type:string;size:255;comment:备注"`
	Perm      string  `gorm:"type:string;size:64;index;comment:权限标识"`
	Type      int     `gorm:"type:tinyint;comment:1:目录 2:菜单 3:按钮"`
	Status    int     `gorm:"type:tinyint;default:0;not null;comment:0: 正常 1: 停用"`
	Visible   int     `gorm:"type:tinyint;default:0;not null;comment:0: 正常 1: 隐藏"`
	Roles     []*Role `gorm:"many2many:role_menu"`
}

func (Menu) TableName() string {
	return "menu"
}
