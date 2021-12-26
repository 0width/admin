package main

import (
	"admin/component/middleware"
	"admin/config"
	"embed"
	"io/fs"
	"net/http"
	"strconv"

	"git.xios.club/xios/gc"
	"github.com/gin-gonic/gin"
)

//go:embed ui/dist
var staticFS embed.FS

//go:generate go run generate/auto_imports/auto_imports.go
func main() {
	server := gin.Default()
	static, err := fs.Sub(staticFS, "ui/dist")
	if err != nil {
		panic(err)
	}
	// 注意： 前端项目 publicPath 也要设置为 /public
	staticRoute := server.Group("/public").Use(middleware.NotModified())
	staticRoute.StaticFS("/", http.FS(static))
	server.GET("/", func(ctx *gin.Context) {
		ctx.Redirect(301, "/public")
	})

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
