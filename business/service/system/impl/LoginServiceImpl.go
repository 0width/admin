package SystemServiceImpl

import (
	systemEntity "admin/business/pogo/entity/system"
	SystemService "admin/business/service/system"
	commonService "admin/common/service"
	"context"
	"encoding/json"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"

	"git.xios.club/xios/gc"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func init() {
	gc.RegisterBean(new(LoginServiceImpl)).Export((*SystemService.LoginService)(nil))
}

type LoginServiceImpl struct {
	JwtService  commonService.JwtService  `autowire:""`
	AuthService commonService.AuthService `autowire:""`
	Key         string                    `value:"${jwt.key}"`
	BufferTime  int64                     `value:"${jwt.bufferTime}"`
	Expire      int64                     `value:"${jwt.expire}"`
	JwtPrefix   string                    `value:"${jwt.prefix}"`
	Db          *gorm.DB                  `autowire:""`
	RedisClient *redis.Client             `autowire:""`
	PermPrefix  string                    `value:"${authFilter.prefix}"`
}

func (this *LoginServiceImpl) Login(userName, password string) (string, error) {
	user := systemEntity.User{}
	if err := this.Db.Where(systemEntity.User{Name: userName}).Find(&user).Error; err != nil {
		return "", err
	}
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", err
	}

	this.AuthService.CachePerms(user.ID)

	cliams := commonService.JwtCliams{
		UserId:   user.ID,
		Username: user.Name,
		NickName: user.NickName,
	}

	cliamsData, _ := json.Marshal(cliams)
	result, err := this.RedisClient.Set(context.Background(), this.JwtPrefix+strconv.Itoa(int(user.ID)), cliamsData,
		time.Duration(this.Expire)*time.Second).Result()
	if err != nil {
		return result, err
	}

	return this.JwtService.CreateToken(cliams, this.Key)
}
