package system

import (
	system2 "admin/business/pogo/dto/system"
	"admin/business/service/system"

	"git.xios.club/xios/gc"
	"github.com/gin-gonic/gin"
)

func init() {
	gc.RegisterBeanFn(func(authRouter *gin.RouterGroup) *MenuController {
		menuController := &MenuController{}
		gr := authRouter.Group("/system/menu")
		{
			gr.GET("/list", menuController.list)
		}
		return menuController
	}, "authRouter")
}

type MenuController struct {
	MenuService system.MenuService `autowire:""`
}

func (this *MenuController) list(ctx *gin.Context) {
	var res []*system2.MenuInfo
	res = this.MenuService.SelectMenuList(ctx.GetUint("userId"))
	ctx.JSON(200, gin.H{
		"code": 200,
		"data": res,
	})
}
