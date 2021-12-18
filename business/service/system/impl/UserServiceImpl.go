package impl

import (
	"admin/business/common"
	common2 "admin/business/pogo/bo/common"
	systemDto "admin/business/pogo/dto/system"
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

func (this *UserServiceImpl) SelectUserList(page *common2.Page) []*systemDto.UserInfo {
	var users []*systemDto.UserInfo
	this.Db.Model(entity.User{}).
		Scopes(common.Paginate(page.Page, page.PageSize)).Find(&users)
	return users
}

func (this *UserServiceImpl) SelectUserById(id uint) *systemDto.UserInfo {
	var user *systemDto.UserInfo
	this.Db.Model(entity.User{}).
		Select("name, nick_name, email, phone, sex, avatar, remark").
		Where("status = 0").Find(&user, id)
	return user
}
