package system

import (
	"git.xios.club/xios/gc"
	"github.com/gin-gonic/gin"
)

func init() {
	gc.RegisterBeanFn(func(authRouter *gin.RouterGroup, g *gin.Engine) *RoleController {
		roleController := &RoleController{}
		groupRoute := authRouter.Group("/system/role")
		groupRoute.GET("/list", roleController.list)
		return roleController
	})
}

type RoleController struct {
}

func (this *RoleController) list(ctx *gin.Context) {

}
