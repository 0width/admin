package system

import (
	"admin/business/common"
	commonBO "admin/business/pogo/bo/common"
	systemBO "admin/business/pogo/bo/system"
	SystemService "admin/business/service/system"

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
			sysUser.POST("/add", userController.addUser)
		}
		return userController
	}, "authRouter")
}

type UserController struct {
	UserService SystemService.UserService `autowire:""`
	Logger      *logrus.Logger            `autowire:""`
}

func (this *UserController) userList(ctx *gin.Context) {
	var page commonBO.Page
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
	res := this.UserService.SelectUserById(ctx.GetUint("userId"))
	ctx.SecureJSON(200, gin.H{
		"code": 200,
		"data": res,
	})
}

func (this *UserController) addUser(ctx *gin.Context) {
	var request systemBO.UserInfo
	if err := ctx.BindJSON(&request); err != nil {
		ctx.JSON(200, gin.H{
			"code": 400,
			"msg":  common.GetError(err, request),
		})
		return
	}
	ctx.JSON(200, "ok")

}
