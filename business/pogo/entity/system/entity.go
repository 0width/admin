package systemEntity

import (
	"git.xios.club/xios/gc"
	"gorm.io/gorm"
)

func init() {
	gc.RegisterBean(new(Entity)).Init(func(entity *Entity) {
		if entity.AutoMigrate {
			_ = entity.Db.AutoMigrate(&SystemMenuEntity{})
			_ = entity.Db.AutoMigrate(&SystemRoleEntity{})
			_ = entity.Db.AutoMigrate(&SystemUserEntity{})
			_ = entity.Db.AutoMigrate(&SystemPostEntity{})
		}
	})
}

type Entity struct {
	Db          *gorm.DB `autowire:""`
	AutoMigrate bool     `value:"${db.autoMigrate:=false}"`
}
