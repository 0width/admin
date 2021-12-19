package system

import (
	"admin/business/pogo/bo/common"
	"admin/business/service/system"

	"github.com/sirupsen/logrus"

	"git.xios.club/xios/gc"
	"github.com/gin-gonic/gin"
)

func init() {
	gc.RegisterBeanFn(func(authRouter *gin.RouterGroup, g *gin.Engine) *UserController {
		userController := &UserController{}
		sysUser := authRouter.Group("/system/user")
		{
			sysUser.GET("/list", userController.userList)
			sysUser.GET("/info", userController.userInfo)
			sysUser.GET("/:id", userController.userInfo)
		}
		return userController
	}, "authRouter")
}

type UserController struct {
	UserService system.UserService `autowire:""`
	Logger      *logrus.Logger     `autowire:""`
}

func (this *UserController) userList(ctx *gin.Context) {
	var page common.Page
	if err := ctx.BindQuery(&page); err != nil {
		ctx.JSON(200, gin.H{
			"code": "401",
			"msg":  "请求参数有误",
		})
		return
	}
	ctx.JSON(200, gin.H{
		"code": 200,
		"data": this.UserService.SelectUserList(&page),
	})
}

func (this *UserController) userInfo(ctx *gin.Context) {
	user := this.UserService.SelectUserById(ctx.GetUint("userId"))
	ctx.SecureJSON(200, gin.H{
		"code": 200,
		"data": user,
	})
}
