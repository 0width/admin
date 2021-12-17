package system

import (
	"admin/business/service/system"

	"git.xios.club/xios/gc"
	"github.com/gin-gonic/gin"
)

func init() {
	gc.RegisterBeanFn(func(authRouter *gin.RouterGroup, g *gin.Engine) *UserController {
		userController := &UserController{}
		sysUser := authRouter.Group("/system/user")
		{
			sysUser.GET("/list", userController.list)
			sysUser.GET("/info", userController.userInfo)
			sysUser.GET("/:id", userController.userInfo)
		}
		return userController
	}, "authRouter")
}

type UserController struct {
	UserService system.UserService `autowire:""`
}

func (this *UserController) list(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"code": 200,
		"data": this.UserService.SelectUserList(),
	})
}

func (this *UserController) userInfo(ctx *gin.Context) {
	user := this.UserService.SelectUserById(ctx.GetUint("userId"))
	ctx.JSON(200, gin.H{
		"code": 200,
		"data": user,
	})
}
