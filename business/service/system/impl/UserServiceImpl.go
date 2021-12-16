package impl

import (
	"admin/business/pogo/entity"
	"admin/business/service/system"

	"git.xios.club/xios/gc"
	"gorm.io/gorm"
)

func init() {
	gc.RegisterBean(new(UserServiceImpl)).Export((*system.UserService)(nil)).Init(func(impl *UserServiceImpl) {

	})
}

type UserServiceImpl struct {
	Db *gorm.DB `autowire:""`
}

func (this *UserServiceImpl) SelectUserList() []entity.User {
	var users []entity.User
	this.Db.Find(&users)
	return users
}
