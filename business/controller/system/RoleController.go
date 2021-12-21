package system

import (
	"admin/business/common"
	systemBO "admin/business/pogo/bo/system"
	SystemService "admin/business/service/system"

	"git.xios.club/xios/gc"
	"github.com/gin-gonic/gin"
)

func init() {
	gc.RegisterBeanFn(func(authRouter *gin.RouterGroup, g *gin.Engine) *RoleController {
		roleController := &RoleController{}
		groupRoute := authRouter.Group("/system/role")
		groupRoute.GET("/list", roleController.list)
		groupRoute.POST("/add", roleController.add)
		groupRoute.PUT("/edit", roleController.edit)
		return roleController
	})
}

type RoleController struct {
	RoleService SystemService.RoleService `autowire:""`
}

func (this *RoleController) list(ctx *gin.Context) {
	res, err := this.RoleService.SelectRoleList(ctx.GetUint("userId"))
	if err != nil {
		ctx.JSON(200, gin.H{
			"code": 500,
			"msg":  err,
		})
	}
	ctx.JSON(200, gin.H{
		"code": 200,
		"data": res,
	})
}

func (this *RoleController) add(ctx *gin.Context) {
	var request systemBO.RoleInfo
	if err := ctx.BindJSON(&request); err != nil {
		ctx.JSON(200, gin.H{
			"code": 400,
			"msg":  common.GetError(err, request),
		})
		return
	}
	if err := this.RoleService.InsertRole(request); err != nil {
		ctx.JSON(200, gin.H{
			"code": 500,
			"msg":  err.Error(),
		})
		return
	}
	ctx.JSON(200, gin.H{
		"code": 200,
		"msg":  "添加成功",
	})
}

func (this *RoleController) edit(ctx *gin.Context) {
	var request systemBO.RoleInfo
	if err := ctx.BindJSON(&request); err != nil {
		common.ValidError(err, request, ctx)
		return
	}
	if err := this.RoleService.UpdateRole(request); err != nil {
		common.InternalError(ctx, err.Error())
		return
	}
	common.SuccessMsg(ctx, "修改成功")
}
