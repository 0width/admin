package authFilter

import (
	"git.xios.club/xios/gc"
	"github.com/gin-gonic/gin"
)

type AuthFilter struct {
	JwtMiddlewre *gin.HandlerFunc `autowire:"jwt"`
}

func init() {
	gc.RegisterBeanFn(func(g *gin.Engine) *AuthFilter {
		authFilter := &AuthFilter{}
		g.Group("/system").Use()
		g.Use(func(ctx *gin.Context) {

		})
		return authFilter
	})
}
