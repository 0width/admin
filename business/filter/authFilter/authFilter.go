package authFilter

import (
	"context"
	"strconv"

	"git.xios.club/xios/gc"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

type AuthFilterConfig struct {
	Prefix              string   `value:"${authFilter.prefix}"`
	SuperManagersConfig []string `value:"${system.superManagers}"`
	SuperManagers       map[string]bool
}

func init() {
	gc.RegisterNameBeanFn("authFilter", func(redisClient *redis.Client, config AuthFilterConfig) gin.HandlerFunc {
		config.SuperManagers = make(map[string]bool)
		for _, v := range config.SuperManagersConfig {
			config.SuperManagers[v] = true
		}
		return func(ctx *gin.Context) {
			if !config.SuperManagers[ctx.GetString("userName")] {
				background := context.Background()
				res, err := redisClient.SIsMember(background, config.Prefix+strconv.Itoa(int(ctx.GetUint("userId"))), ctx.FullPath()).Result()
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
			}
			ctx.Next()
		}
	})
}
