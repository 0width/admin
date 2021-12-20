package system

import (
	systemDTO "admin/business/pogo/dto/system"
	SystemService "admin/business/service/system"

	"git.xios.club/xios/gc"
	"github.com/gin-gonic/gin"
)

func init() {
	gc.RegisterBeanFn(func(authRouter *gin.RouterGroup) *SystemMenuController {
		menuController := &SystemMenuController{}
		gr := authRouter.Group("/system/menu")
		{
			gr.GET("/list", menuController.list)
		}
		return menuController
	}, "authRouter")
}

type SystemMenuController struct {
	MenuService SystemService.SystemMenuService `autowire:""`
}

func (this *SystemMenuController) list(ctx *gin.Context) {
	var res []*systemDTO.SystemMenuInfoDTO
	res = this.MenuService.SelectMenuList(ctx.GetUint("userId"))
	ctx.JSON(200, gin.H{
		"code": 200,
		"data": res,
	})
}
