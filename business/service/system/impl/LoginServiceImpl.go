package impl

import (
	"admin/business/pogo/entity"
	"admin/business/service/common"
	"admin/business/service/system"
	"time"

	"github.com/go-redis/redis"

	"git.xios.club/xios/gc"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func init() {
	gc.RegisterBean(new(LoginServiceImpl)).Export((*system.LoginService)(nil))
}

type LoginServiceImpl struct {
	JwtService  common.JwtService  `autowire:""`
	AuthService common.AuthService `autowire:""`
	Key         string             `value:"${jwt.key}"`
	BufferTime  int64              `value:"${jwt.bufferTime}"`
	Expire      int64              `value:"${jwt.expire}"`
	Db          *gorm.DB           `autowire:""`
	RedisClient *redis.Client      `autowire:""`
	PermPrefix  string             `value:"${authFilter.prefix}"`
}

func (this *LoginServiceImpl) Login(userName, password string) (string, error) {
	user := entity.User{}
	if err := this.Db.Where(entity.User{Name: userName}).Find(&user).Error; err != nil {
		return "", err
	}
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", err
	}

	this.AuthService.CachePerms(user.ID)

	return this.JwtService.CreateToken(common.JwtCliams{
		UserId:     user.ID,
		Username:   user.Name,
		NickName:   user.NickName,
		BufferTime: this.BufferTime,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Unix() + this.Expire,
		},
	}, this.Key)
}
