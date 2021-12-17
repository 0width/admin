package common

import (
	"github.com/dgrijalva/jwt-go"
	uuid "github.com/satori/go.uuid"
)

type JwtCustomClaims struct {
	UUID        uuid.UUID
	ID          uint
	Username    string
	NickName    string
	AuthorityId string
	BufferTime  int64
	jwt.StandardClaims
}

type JwtService interface {
	CreateToken(claims JwtCustomClaims, key string) (string, error)
	CreateTokenByOldToken(oldToken, key string, claims JwtCustomClaims) (string, error)
	ParseToken(tokenString, key string) (*JwtCustomClaims, error)
}
