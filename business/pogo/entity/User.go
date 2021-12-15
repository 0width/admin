package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model

}

func (User) TableName() string {
	return "user"
}