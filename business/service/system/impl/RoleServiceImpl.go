package SystemServiceImpl

import (
	systemBO "admin/business/pogo/bo/system"
	systemDTO "admin/business/pogo/dto/system"
	systemEntity "admin/business/pogo/entity/system"
	SystemService "admin/business/service/system"

	"git.xios.club/xios/gc"
	"gorm.io/gorm"
)

func init() {
	gc.RegisterBean(new(RoleServiceImpl)).Export((*SystemService.RoleService)(nil))
}

type RoleServiceImpl struct {
	Db *gorm.DB `autowire:""`
}

func (this *RoleServiceImpl) SelectRoleList(userId uint) ([]systemDTO.RoleInfo, error) {
	var roles []systemDTO.RoleInfo
	if err := this.Db.Model(systemEntity.User{}).Association("Role").Find(roles, userId); err != nil {
		return nil, err
	}
	return roles, nil
}

func (this *RoleServiceImpl) InsertRole(info systemBO.RoleInfo) error {
	role := systemEntity.Role{
		Name:   info.Name,
		Sort:   info.Sort,
		Status: 0,
		Remark: info.Remark,
	}
	return this.Db.Create(role).Error
}

func (this *RoleServiceImpl) UpdateRole(info systemBO.RoleInfo) error {

}

func (this *RoleServiceImpl) insertMenuByRole(info systemBO.RoleInfo) {

}
