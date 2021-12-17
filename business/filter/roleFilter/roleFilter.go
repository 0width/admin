package roleFilter

import (
	"strconv"

	"git.xios.club/xios/gc"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

type RoleFilterConfig struct {
	Prefix string `value:"${roleFilter.prefix}"`
}

func init() {
	gc.RegisterNameBeanFn("roleFilter", func(redisClient *redis.Client, config RoleFilterConfig) gin.HandlerFunc {
		return func(ctx *gin.Context) {
			res, err := redisClient.SIsMember(config.Prefix+strconv.Itoa(int(ctx.GetUint("userId"))), ctx.FullPath()).Result()
			if err != nil {
				ctx.AbortWithStatusJSON(200, gin.H{
					"code": 500,
					"msg":  "内部服务异常: 0c371256-5f34-11ec-887f-525400354c67",
				})
				return
			}
			if !res {
				ctx.AbortWithStatusJSON(200, gin.H{
					"code": 401,
					"msg":  "权限不足",
				})
				return
			}
			ctx.Next()
		}
	})
}
