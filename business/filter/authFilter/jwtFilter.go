package authFilter

import (
	"admin/business/common/constant"
	commonService "admin/business/service/common"
	commonServiceImpl "admin/business/service/common/impl"
	"context"
	"strconv"

	"git.xios.club/xios/gc"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

func init() {
	gc.RegisterNameBeanFn("jwt",
		func(jwtService commonService.JwtService, authService commonService.AuthService,
			jwtConfig JwtConfig, redisClient *redis.Client) gin.HandlerFunc {

			return func(ctx *gin.Context) {
				token := ctx.Request.Header.Get("x-token")
				if token == "" {
					ctx.JSON(200, gin.H{
						"code": 401,
						"msg":  "未登录或非法访问",
					})
					ctx.Abort()
					return
				}
				claims, err := jwtService.ParseToken(token, jwtConfig.Key)
				if err != nil {
					if err == commonServiceImpl.TokenExpired {
						ctx.JSON(200, gin.H{
							"code": 401,
							"msg":  "授权已过期",
						})
						ctx.Abort()
						return
					}
					ctx.JSON(200, gin.H{
						"code": 401,
						"msg":  err.Error(),
					})
					ctx.Abort()
					return
				}

				_, err = redisClient.Get(context.Background(), jwtConfig.Prefix+strconv.Itoa(int(claims.UserId))).Result()
				if err == redis.Nil {
					ctx.JSON(200, gin.H{
						"code": 401,
						"msg":  "用户未登录或授权已过期",
					})
					ctx.Abort()
					return
				} else if err != nil {
					ctx.JSON(200, gin.H{
						"code": 500,
						"msg":  "内部服务异常: f42890e4-7285-11ec-88ee-525400354c67",
					})
					ctx.Abort()
					return
				}

				ctx.Set("userId", claims.UserId)
				ctx.Set("userName", claims.Username)
				ctx.Set(constant.CLAIMS, claims)

				ctx.Next()
			}
		})
}

type JwtConfig struct {
	Key    string `value:"${jwt.key}"`
	Expire int64  `value:"${jwt.expire}"`
	Prefix string `value:"${jwt.prefix}"`
}
