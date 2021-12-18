package controller

import (
	"git.xios.club/xios/gc"
	"github.com/gin-gonic/gin"
)

func init() {
	gc.RegisterNameBeanFn("authRouter", func(jwtMiddleware gin.HandlerFunc, roleFilter gin.HandlerFunc, g *gin.Engine) *gin.RouterGroup {
		group := g.Group("/")
		group.Use(jwtMiddleware, roleFilter)
		return group
	}, "jwt", "roleFilter")
}
