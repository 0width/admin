package system

import (
	systemDTO "admin/business/pogo/dto/system"
	SystemService "admin/business/service/system"

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
	MenuService SystemService.SystemMenuService `autowire:""`
}

func (this *MenuController) list(ctx *gin.Context) {
	var res []*systemDTO.SystemMenuInfoDTO
	res = this.MenuService.SelectMenuList(ctx.GetUint("userId"))
	ctx.JSON(200, gin.H{
		"code": 200,
		"data": res,
	})
}
