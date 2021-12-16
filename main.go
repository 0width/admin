package main

import (
	_ "admin/business/pogo/entity"
	_ "admin/component/db/mysql"

	_ "admin/business/controller"
	_ "admin/business/service"

	"git.xios.club/xios/gc"
	"github.com/gin-gonic/gin"
)

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
