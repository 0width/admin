package system

import (
	systemBO "admin/business/pogo/bo/system"
	systemDTO "admin/business/pogo/dto/system"
	systemService "admin/business/service/system"
	"admin/common"
	"strconv"

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
			gr.GET("/:menuId", menuController.get)
			gr.DELETE("/:menuId", menuController.delete)
		}
		return menuController
	}, "authRouter")
}

type MenuController struct {
	MenuService systemService.MenuService `autowire:""`
}

func (this *MenuController) list(ctx *gin.Context) {
	var res []*systemDTO.MenuInfo
	res = this.MenuService.SelectMenuList()
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

func (this *MenuController) get(ctx *gin.Context) {
	menuId, err := strconv.Atoi(ctx.Param("menuId"))
	if err != nil {
		ctx.JSON(200, gin.H{
			"code": 400,
			"msg":  "参数数据类型错误",
		})
	}
	if resp, err := this.MenuService.SelectMenuById(uint(menuId)); err != nil {
		common.InternalError(ctx, err.Error())
		return
	} else {
		common.SuccessData(ctx, resp)
	}
}

func (this *MenuController) delete(ctx *gin.Context) {
	menuId, err := strconv.Atoi(ctx.Param("menuId"))
	if err != nil {
		ctx.JSON(200, gin.H{
			"code": 400,
			"msg":  "参数数据类型错误",
		})
	}
	if err := this.MenuService.DeleteMenuById(uint(menuId)); err != nil {
		common.InternalError(ctx, err.Error())
		return
	}
	common.SuccessMsg(ctx, "删除成功")
}
