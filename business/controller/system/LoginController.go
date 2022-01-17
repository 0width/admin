package system

import (
	systemBO "admin/business/pogo/bo/system"
	systemService "admin/business/service/system"
	"admin/common"

	"git.xios.club/xios/gc"
	"github.com/gin-gonic/gin"
)

func init() {
	gc.RegisterBeanFn(func(authRouter *gin.RouterGroup, g *gin.Engine) *LoginController {
		l := &LoginController{}
		g.POST("/system/user/login", l.login)
		authRouter.POST("/system/user/logout", l.logout)
		return l
	}, "authRouter")
}

type LoginController struct {
	LoginService systemService.LoginService `autowire:""`
}

func (this *LoginController) login(ctx *gin.Context) {
	request := systemBO.Login{}
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(200, gin.H{
			"code": 400,
			"msg":  common.GetError(err, request),
		})
		return
	}
	token, err := this.LoginService.Login(request.UserName, request.Password)
	if err != nil {
		ctx.JSON(200, gin.H{
			"code": 401,
			"msg":  "用户名或密码错误",
		})
		return
	}
	ctx.JSON(200, gin.H{
		"code": 200,
		"data": gin.H{
			"token": token,
		},
	})
}

func (this *LoginController) logout(ctx *gin.Context) {
	_ = this.LoginService.Logout(ctx.GetUint("userId"))
	ctx.JSON(200, gin.H{
		"code": 200,
		"msg":  "登出成功",
	})
}
