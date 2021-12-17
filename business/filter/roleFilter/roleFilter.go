package roleFilter

import (
	"fmt"

	"git.xios.club/xios/gc"
	"github.com/gin-gonic/gin"
)

func init() {
	gc.RegisterNameBeanFn("roleFilter", func() gin.HandlerFunc {
		return func(ctx *gin.Context) {
			fmt.Println(ctx.GetUint("userId"))
		}
	})
}
