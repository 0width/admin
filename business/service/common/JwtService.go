package common

import (
	"github.com/dgrijalva/jwt-go"
)

type JwtCustomClaims struct {
	UserId     uint
	Username   string
	NickName   string
	BufferTime int64
	jwt.StandardClaims
}

type JwtService interface {
	CreateToken(claims JwtCustomClaims, key string) (string, error)
	CreateTokenByOldToken(oldToken, key string, claims JwtCustomClaims) (string, error)
	ParseToken(tokenString, key string) (*JwtCustomClaims, error)
}
