package SystemServiceImpl

import (
	systemEntity "admin/business/pogo/entity/system"
	commonService "admin/business/service/common"
	SystemService "admin/business/service/system"
	"time"

	"github.com/go-redis/redis"

	"git.xios.club/xios/gc"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func init() {
	gc.RegisterBean(new(SystemLoginServiceImpl)).Export((*SystemService.SystemLoginService)(nil))
}

type SystemLoginServiceImpl struct {
	JwtService  commonService.CommonJwtService  `autowire:""`
	AuthService commonService.CommonAuthService `autowire:""`
	Key         string                          `value:"${jwt.key}"`
	BufferTime  int64                           `value:"${jwt.bufferTime}"`
	Expire      int64                           `value:"${jwt.expire}"`
	Db          *gorm.DB                        `autowire:""`
	RedisClient *redis.Client                   `autowire:""`
	PermPrefix  string                          `value:"${authFilter.prefix}"`
}

func (this *SystemLoginServiceImpl) Login(userName, password string) (string, error) {
	user := systemEntity.SystemUserEntity{}
	if err := this.Db.Where(systemEntity.SystemUserEntity{Name: userName}).Find(&user).Error; err != nil {
		return "", err
	}
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", err
	}

	this.AuthService.CachePerms(user.ID)

	return this.JwtService.CreateToken(commonService.JwtCliams{
		UserId:     user.ID,
		Username:   user.Name,
		NickName:   user.NickName,
		BufferTime: this.BufferTime,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Unix() + this.Expire,
		},
	}, this.Key)
}
