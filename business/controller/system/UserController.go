package system

import (
	"admin/business/service/system"
	"strconv"

	_ "admin/component/jwt"

	"git.xios.club/xios/gc"
	"github.com/gin-gonic/gin"
)

func init() {
	gc.RegisterBeanFn(func(jwtMiddleware gin.HandlerFunc, g *gin.Engine) *UserController {
		userController := &UserController{}
		userGroup := g.Group("/system/user").Use(jwtMiddleware)
		{
			userGroup.GET("/list", userController.list)
			userGroup.GET("/info", userController.userInfo)
		}
		return userController
	}, "jwt")
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
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(200, gin.H{
			"code": 400,
			"msg":  "ID参数类型错误",
		})
		return
	}
	ctx.JSON(200, gin.H{
		"code": 200,
		"data": id,
	})
}
