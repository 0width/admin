package SystemServiceImpl

import (
	"admin/business/common"
	commonBO "admin/business/pogo/bo/common"
	systemDTO "admin/business/pogo/dto/system"
	systemEntity "admin/business/pogo/entity/system"
	SystemService "admin/business/service/system"

	"git.xios.club/xios/gc"
	"gorm.io/gorm"
)

func init() {
	gc.RegisterBean(new(UserServiceImpl)).Export((*SystemService.UserService)(nil)).Init(func(impl *UserServiceImpl) {

	})
}

type UserServiceImpl struct {
	Db *gorm.DB `autowire:""`
}

func (this *UserServiceImpl) SelectUserList(page *commonBO.Page) []*systemDTO.UserInfo {
	var users []*systemDTO.UserInfo
	this.Db.Model(systemEntity.User{}).
		Scopes(common.Paginate(page.Page, page.PageSize)).Find(&users)
	return users
}

func (this *UserServiceImpl) SelectUserById(id uint) *systemDTO.UserInfo {
	var user *systemDTO.UserInfo
	this.Db.Model(systemEntity.User{}).
		Select("id, name, nick_name, email, phone, sex, avatar, remark").
		Where("status = 0").Find(&user, id)
	return user
}
