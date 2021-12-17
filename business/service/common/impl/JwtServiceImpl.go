package impl

import (
	"admin/business/service/common"
	"errors"

	"git.xios.club/xios/gc"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/sync/singleflight"
)

func init() {
	gc.RegisterBean(new(JwtServiceImpl)).Export((*common.JwtService)(nil))
}

var (
	TokenExpired     = errors.New("Token is expired")
	TokenNotValidYet = errors.New("Token not active yet")
	TokenMalformed   = errors.New("That's not even a token")
	TokenInvalid     = errors.New("Couldn't handle this token")
)

type JwtServiceImpl struct {
}

func (JwtServiceImpl) CreateToken(claims common.JwtCustomClaims, key string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(key)
}

func (this *JwtServiceImpl) CreateTokenByOldToken(oldToken, key string, claims common.JwtCustomClaims) (string, error) {
	group := &singleflight.Group{}
	v, err, _ := group.Do("JWT:"+oldToken, func() (interface{}, error) {
		return this.CreateToken(claims, key)
	})
	return v.(string), err
}

func (this *JwtServiceImpl) ParseToken(tokenString, key string) (*common.JwtCustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &common.JwtCustomClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return []byte(key), nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}
	if token != nil {
		if claims, ok := token.Claims.(*common.JwtCustomClaims); ok && token.Valid {
			return claims, nil
		}
		return nil, TokenInvalid

	} else {
		return nil, TokenInvalid
	}
}
