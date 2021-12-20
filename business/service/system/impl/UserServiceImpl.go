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
	gc.RegisterBean(new(SystemUserServiceImpl)).Export((*SystemService.SystemUserService)(nil)).Init(func(impl *SystemUserServiceImpl) {

	})
}

type SystemUserServiceImpl struct {
	Db *gorm.DB `autowire:""`
}

func (this *SystemUserServiceImpl) SelectUserList(page *commonBO.CommonPage) []*systemDTO.SystemUserInfoDTO {
	var users []*systemDTO.SystemUserInfoDTO
	this.Db.Model(systemEntity.SystemUserEntity{}).
		Scopes(common.Paginate(page.Page, page.PageSize)).Find(&users)
	return users
}

func (this *SystemUserServiceImpl) SelectUserById(id uint) *systemDTO.SystemUserInfoDTO {
	var user *systemDTO.SystemUserInfoDTO
	this.Db.Model(systemEntity.SystemUserEntity{}).
		Select("id, name, nick_name, email, phone, sex, avatar, remark").
		Where("status = 0").Find(&user, id)
	return user
}
