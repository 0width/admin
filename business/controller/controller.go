package controller

import (
	"git.xios.club/xios/gc"
	"github.com/gin-gonic/gin"
)

func init() {
	gc.RegisterNameBeanFn("authRouter", func(jwtMiddleware gin.HandlerFunc, authFilter gin.HandlerFunc, g *gin.Engine) *gin.RouterGroup {
		group := g.Group("/")
		group.Use(jwtMiddleware, authFilter)
		return group
	}, "jwt", "authFilter")
}
