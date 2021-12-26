package main

import (
	"admin/config"
	"strconv"

	"git.xios.club/xios/gc"
	"github.com/gin-gonic/gin"
)

//go:generate go run generate/auto_imports/auto_imports.go
//go:generate go run generate/migrate/migrate.go
func main() {
	server := gin.Default()

	gc.RegisterBean(server)

	confDir := "./config"

	gc.LoadProperties(confDir + "/application.yml")
	if confActive := gc.GetStringProperty("active"); confActive != "" {
		gc.LoadProperties(confDir + "/application-" + confActive + ".yml")
	}
	gc.AutoWireBeans()

	var webConfig *config.WebConfig
	gc.GetBean(&webConfig)

	_ = server.Run(":" + strconv.Itoa(webConfig.Port))
}
