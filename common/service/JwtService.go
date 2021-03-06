package commonService

import (
	"github.com/dgrijalva/jwt-go"
)

type JwtCliams struct {
	UserId   uint
	Username string
	NickName string
	jwt.StandardClaims
}

type JwtService interface {
	CreateToken(claims JwtCliams, key string) (string, error)
	CreateTokenByOldToken(oldToken, key string, claims JwtCliams) (string, error)
	ParseToken(tokenString, key string) (*JwtCliams, error)
}
