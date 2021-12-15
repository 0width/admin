package entity

import "gorm.io/gorm"

type Role struct {
	gorm.Model

}

func (Role) TableName() string {
	return "role"
}