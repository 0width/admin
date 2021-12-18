package main

import (
	"git.xios.club/xios/gc"
	"github.com/gin-gonic/gin"
)

//go:generate go run auto_imports/auto_imports.go
func main() {
	server := gin.Default()

	gc.RegisterBean(server)

	confDir := "./config"

	gc.LoadProperties(confDir + "/application.yml")
	if confActive := gc.GetStringProperty("active"); confActive != "" {
		gc.LoadProperties(confDir + "/application-" + confActive + ".yml")
	}
	gc.AutoWireBeans()

	_ = server.Run(":8080")
}
