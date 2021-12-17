package entity

import "gorm.io/gorm"

// Role 角色表
type Role struct {
	gorm.Model
	Name   string  `gorm:"size:255"`
	Sort   int     `gorm:"type:int;size:32;comment:显示顺序"`
	Status int     `gorm:"type:tinyint;default:0;comment:0: 正常 1: 停用"`
	Remark string  `gorm:"size:1023;comment:备注"`
	Users  []*User `gorm:"many2many:user_role"`
	Menus  []*Menu `gorm:"many2many:role_menu"`
}

func (Role) TableName() string {
	return "role"
}
