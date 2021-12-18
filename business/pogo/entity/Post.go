package entity

import "gorm.io/gorm"

// Post 岗位表
type Post struct {
	gorm.Model
}

func (Post) TableName() string {
	return "post"
}
