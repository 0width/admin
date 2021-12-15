package main

import (
	"git.xios.club/xios/gc"
	"github.com/gin-gonic/gin"

	_ "admin/business/pogo/entity"
	_ "admin/component/db/mysql"
)

func main() {
	server := gin.Default()

	gc.RegisterBean(&server)
	gc.LoadProperties("./config/application.yml")
	gc.AutoWireBeans()

	_ = server.Run(":8080")
}
