package SystemServiceImpl

import (
	systemBO "admin/business/pogo/bo/system"
	systemDTO "admin/business/pogo/dto/system"
	systemEntity "admin/business/pogo/entity/system"
	SystemService "admin/business/service/system"
	"errors"

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

func (this *RoleServiceImpl) InsertRole(roleInfo systemBO.RoleInfo) error {
	if this.hasRoleWithName(roleInfo.Name) {
		return errors.New("角色名已经存在")
	}
	role := systemEntity.Role{
		Name:   roleInfo.Name,
		Sort:   roleInfo.Sort,
		Remark: roleInfo.Remark,
	}
	if err := this.Db.Create(&role).Error; err != nil {
		return err
	}
	return this.insertMenuByRole(role.ID, roleInfo.MenuIds)
}

func (this *RoleServiceImpl) UpdateRole(roleInfo systemBO.RoleInfo) error {
	role := &systemEntity.Role{
		Model:  gorm.Model{ID: roleInfo.ID},
		Name:   roleInfo.Name,
		Sort:   roleInfo.Sort,
		Remark: roleInfo.Remark,
	}
	if err := this.Db.Updates(&role).Error; err != nil {
		return err
	}
	if err := this.Db.Model(&role).
		Association("Menus").Clear(); err != nil {
		return err
	}
	return this.insertMenuByRole(roleInfo.ID, roleInfo.MenuIds)
}

func (this *RoleServiceImpl) insertMenuByRole(roleId uint, menuIds []uint) error {
	datas := []map[string]interface{}{}
	for _, v := range menuIds {
		datas = append(datas, map[string]interface{}{"role_id": roleId, "menu_id": v})
	}
	return this.Db.Table("role_menu").Create(datas).Error
}

func (this *RoleServiceImpl) hasRoleWithName(name string) bool {
	r := this.Db.First(&systemEntity.Role{Name: name})
	if r.Error != nil {
		return false
	}
	return true
}
