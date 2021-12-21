package system

import (
	"admin/business/common"
	systemBO "admin/business/pogo/bo/system"
	SystemService "admin/business/service/system"

	"git.xios.club/xios/gc"
	"github.com/gin-gonic/gin"
)

func init() {
	gc.RegisterBeanFn(func(g *gin.Engine) *LoginController {
		l := &LoginController{}
		g.POST("/system/user/login", l.login)
		return l
	})
}

type LoginController struct {
	LoginService SystemService.LoginService `autowire:""`
}

func (this *LoginController) login(ctx *gin.Context) {
	request := systemBO.Login{}
	if err := ctx.BindJSON(&request); err != nil {
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
