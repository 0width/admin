package systemEntity

import (
	"context"
	"os"
	"strings"
	"time"

	"git.xios.club/xios/gc"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func init() {
	gc.RegisterBeanFn(func() *Migrator {
		return &Migrator{Interface: logger.Default.LogMode(logger.Info)}
	}).Init(func(migrator *Migrator) {
		if migrator.AutoMigrate {
			err := migrator.migrate(&Dept{}, &Menu{}, &Role{}, &Post{})
			if err != nil {
				panic(err)
			}
			migrator.printOut()
		}
	})
}

type Migrator struct {
	logger.Interface
	Statements  []string
	Db          *gorm.DB `autowire:""`
	AutoMigrate bool     `value:"${db.autoMigrate:=false}"`
	MigrateFile string   `value:"${db.migrateFile:=}"`
}

func (this *Migrator) migrate(dst ...interface{}) error {
	session := this.Db.Session(&gorm.Session{
		Logger: this,
	})
	return session.AutoMigrate(dst...)
}

func (this *Migrator) printOut() {
	if len(this.Statements) > 0 {
		currentTime := time.Now().Format("2006-01-02-15-04-05")
		sqls := ""
		for _, v := range this.Statements {
			sqls += v + ";\n"
		}
		if this.MigrateFile != "" {
			migrateFile, err := os.Create(this.MigrateFile + "." + currentTime + ".sql")
			if err != nil {
				panic(err)
			}
			defer migrateFile.Close()

			_, err = migrateFile.Write([]byte(sqls))
			if err != nil {
				panic(err)
			}
		}
		print(sqls)
	}
}

func (this *Migrator) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	sql, _ := fc()
	if strings.HasPrefix(sql, "CREATE") || strings.HasPrefix(sql, "DROP") || strings.HasPrefix(sql, "ALTER") {
		this.Statements = append(this.Statements, sql)
	}
}
