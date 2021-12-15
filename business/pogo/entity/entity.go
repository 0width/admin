package entity

import (
	"git.xios.club/xios/gc"
	"gorm.io/gorm"
)

func init() {
	gc.RegisterBean(new(Entity)).Init(func(entity *Entity) {
		if entity.AutoMigrate {
			_ = entity.Db.AutoMigrate(&Menu{})
			_ = entity.Db.AutoMigrate(&Role{})
			_ = entity.Db.AutoMigrate(&User{})
		}
	})
}

type Entity struct {
	Db *gorm.DB `autowire:""`
	AutoMigrate bool `value:"${db.autoMigrate:=false}"`
}