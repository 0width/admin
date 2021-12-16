package db

type DbConfig struct {
	Url                                      string `value:"${db.url}"`
	Debug                                    bool   `value:"${db.debug:=false}"`
	AutoMigrate                              bool   `value:"${db.autoMigrate}"`
	DisableForeignKeyConstraintWhenMigrating bool   `value:"${db.disableForeignKeyConstraintWhenMigrating:=false}"`
}
