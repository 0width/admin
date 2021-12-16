package system

import (
	"admin/business/service/system"

	"git.xios.club/xios/gc"
	"github.com/gin-gonic/gin"
)

func init() {
	gc.RegisterBeanFn(func(g *gin.Engine) *UserController {
		userController := &UserController{}
		g.GET("/system/user/list", userController.list)
		return userController
	})
}

type UserController struct {
	UserService system.UserService `autowire:""`
}

func (this UserController) list(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"code": 200,
		"data": this.UserService.SelectUserList(),
	})
}
