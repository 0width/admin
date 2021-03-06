package commonServiceImpl

import (
	commonService "admin/common/service"
	"context"
	"strconv"
	"time"

	"git.xios.club/xios/gc"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

func init() {
	gc.RegisterBean(new(AuthServiceImpl)).Export((*commonService.AuthService)(nil))
}

type AuthServiceImpl struct {
	Db          *gorm.DB      `autowire:""`
	PermPrefix  string        `value:"${authFilter.prefix}"`
	Expire      int           `value:"${jwt.expire}"`
	RedisClient *redis.Client `autowire:""`
}

func (this *AuthServiceImpl) CachePerms(userId uint) {
	perms := []string{}
	this.Db.Table("user_role a").
		Joins("left join role_menu b on a.role_id = b.role_id").
		Joins("left join menu c on b.menu_id = c.id").
		Where("a.user_id = ?", userId).Where("c.status = 0").
		Group("c.perm").Pluck("c.perm", &perms)

	key := this.PermPrefix + strconv.Itoa(int(userId))
	background := context.Background()
	this.RedisClient.SAdd(background, key, perms)
	this.RedisClient.Expire(background, key, time.Duration(this.Expire)*time.Second)
}
