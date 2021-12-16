package mysql

import (
	"admin/component/db"

	"git.xios.club/xios/gc"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func init() {
	gc.RegisterBeanFn(func(config db.DbConfig) (*gorm.DB, error) {
		var gormConfig gorm.Config
		if config.Debug {
			gormConfig.Logger = logger.Default.LogMode(logger.Info)
		}
		gormConfig.DisableForeignKeyConstraintWhenMigrating = config.DisableForeignKeyConstraintWhenMigrating
		return gorm.Open(mysql.Open(config.Url), &gormConfig)
	})
}
