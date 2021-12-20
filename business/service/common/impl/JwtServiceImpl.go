package commonServiceImpl

import (
	commonService "admin/business/service/common"
	"errors"

	"git.xios.club/xios/gc"

	"github.com/golang-jwt/jwt"
	"golang.org/x/sync/singleflight"
)

func init() {
	gc.RegisterBean(new(CommonJwtServiceImpl)).Export((*commonService.CommonJwtService)(nil))
}

var (
	TokenExpired     = errors.New("Token is expired")
	TokenNotValidYet = errors.New("Token not active yet")
	TokenMalformed   = errors.New("That's not even a token")
	TokenInvalid     = errors.New("Couldn't handle this token")
)

type CommonJwtServiceImpl struct {
}

func (CommonJwtServiceImpl) CreateToken(claims commonService.JwtCliams, key string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(key))
}

func (this *CommonJwtServiceImpl) CreateTokenByOldToken(oldToken, key string, claims commonService.JwtCliams) (string, error) {
	group := &singleflight.Group{}
	v, err, _ := group.Do("JWT:"+oldToken, func() (interface{}, error) {
		return this.CreateToken(claims, key)
	})
	return v.(string), err
}

func (this *CommonJwtServiceImpl) ParseToken(tokenString, key string) (*commonService.JwtCliams, error) {
	token, err := jwt.ParseWithClaims(tokenString, &commonService.JwtCliams{}, func(token *jwt.Token) (i interface{}, e error) {
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
		if claims, ok := token.Claims.(*commonService.JwtCliams); ok && token.Valid {
			return claims, nil
		}
		return nil, TokenInvalid

	} else {
		return nil, TokenInvalid
	}
}
