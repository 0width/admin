package jwt

import (
	"admin/business/common/constant"
	commonService "admin/business/service/common"
	commonServiceImpl "admin/business/service/common/impl"
	"strconv"
	"time"

	"git.xios.club/xios/gc"
	"github.com/gin-gonic/gin"
)

func init() {
	gc.RegisterNameBeanFn("jwt", func(jwtService commonService.CommonJwtService, authService commonService.CommonAuthService, jwtConfig JwtConfig) gin.HandlerFunc {
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

			ctx.Set("userId", claims.UserId)
			ctx.Set("userName", claims.Username)
			ctx.Set(constant.CLAIMS, claims)

			// 续期
			if claims.ExpiresAt-time.Now().Unix() < claims.BufferTime {
				claims.ExpiresAt = time.Now().Unix() + jwtConfig.Expire
				newToken, _ := jwtService.CreateTokenByOldToken(token, jwtConfig.Key, *claims)
				newClaims, _ := jwtService.ParseToken(newToken, jwtConfig.Key)
				ctx.Header("new-token", newToken)
				ctx.Header("new-expires-at", strconv.FormatInt(newClaims.ExpiresAt, 10))

				authService.CachePerms(claims.UserId)
			}

			ctx.Next()
		}
	})
}

type JwtConfig struct {
	Key    string `value:"${jwt.key}"`
	Expire int64  `value:"${jwt.expire}"`
}
