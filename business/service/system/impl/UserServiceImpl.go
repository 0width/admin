package SystemServiceImpl

import (
	"admin/business/common"
	commonBO "admin/business/pogo/bo/common"
	systemBO "admin/business/pogo/bo/system"
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

func (this *UserServiceImpl) UpdateRoles(userId uint, roleIds ...uint) error {
	if err := this.Db.Model(&systemEntity.User{Model: gorm.Model{ID: userId}}).
		Association("Roles").Clear(); err != nil {
		return err
	}
	datas := []map[string]interface{}{}
	for _, v := range roleIds {
		datas = append(datas, map[string]interface{}{"user_id": userId, "role_id": v})
	}
	if len(datas) == 0 {
		return nil
	}
	return this.Db.Table("user_role").Create(datas).Error
}

func (this *UserServiceImpl) UpdateUser(userInfo systemBO.UpdateUserInfo) error {
	user := systemEntity.User{
		Model: gorm.Model{
			ID: userInfo.ID,
		},
		NickName: userInfo.NickName,
		Email:    userInfo.Email,
		Phone:    userInfo.Phone,
		Sex:      userInfo.Sex,
		Remark:   userInfo.Remark,
	}
	if err := this.Db.Updates(user).Error; err != nil {
		return err
	}
	return this.UpdateRoles(userInfo.ID, userInfo.RoleIds...)
}
