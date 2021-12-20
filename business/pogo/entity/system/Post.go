package systemEntity

import "gorm.io/gorm"

// SystemPostEntity 岗位表
type SystemPostEntity struct {
	gorm.Model
}

func (SystemPostEntity) TableName() string {
	return "post"
}
