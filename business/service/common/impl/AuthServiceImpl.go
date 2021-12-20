package commonServiceImpl

import (
	commonService "admin/business/service/common"
	"strconv"
	"time"

	"git.xios.club/xios/gc"
	"github.com/go-redis/redis"
	"gorm.io/gorm"
)

func init() {
	gc.RegisterBean(new(CommonAuthServiceImpl)).Export((*commonService.CommonAuthService)(nil))
}

type CommonAuthServiceImpl struct {
	Db          *gorm.DB      `autowire:""`
	PermPrefix  string        `value:"${authFilter.prefix}"`
	Expire      int           `value:"${jwt.expire}"`
	RedisClient *redis.Client `autowire:""`
}

func (this *CommonAuthServiceImpl) CachePerms(userId uint) {
	perms := []string{}
	this.Db.Table("user_role a").
		Joins("left join role_menu b on a.role_id = b.role_id").
		Joins("left join menu c on b.menu_id = c.id").
		Where("a.user_id = ?", userId).Where("c.status = 0").
		Group("c.perm").Pluck("c.perm", &perms)

	key := this.PermPrefix + strconv.Itoa(int(userId))
	this.RedisClient.SAdd(key, perms)
	this.RedisClient.Expire(key, time.Duration(this.Expire)*time.Second)
}
