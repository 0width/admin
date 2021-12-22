package system

import (
	"admin/business/common"
	systemBO "admin/business/pogo/bo/system"
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
			gr.POST("/add", menuController.add)
			gr.PUT("/edit", menuController.edit)
		}
		return menuController
	}, "authRouter")
}

type MenuController struct {
	MenuService SystemService.MenuService `autowire:""`
}

func (this *MenuController) list(ctx *gin.Context) {
	var res []*systemDTO.MenuInfo
	res = this.MenuService.SelectMenuList(ctx.GetUint("userId"))
	ctx.JSON(200, gin.H{
		"code": 200,
		"data": res,
	})
}

func (this *MenuController) add(ctx *gin.Context) {
	var request systemBO.AddMenuInfo
	if err := ctx.ShouldBindJSON(&request); err != nil {
		common.ValidError(err, request, ctx)
		return
	}
	if err := this.MenuService.InsertMenu(request); err != nil {
		common.InternalError(ctx, err.Error())
		return
	}
	common.SuccessMsg(ctx, "添加成功")
}

func (this *MenuController) edit(ctx *gin.Context) {
	var request systemBO.EditMenuInfo
	if err := ctx.ShouldBindJSON(&request); err != nil {
		common.ValidError(err, request, ctx)
		return
	}
	if err := this.MenuService.UpdateMenu(request); err != nil {
		common.InternalError(ctx, err.Error())
		return
	}
	common.SuccessMsg(ctx, "修改成功")
}
