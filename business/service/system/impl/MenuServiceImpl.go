package systemServiceImpl

import (
	systemBO "admin/business/pogo/bo/system"
	systemDTO "admin/business/pogo/dto/system"
	systemEntity "admin/business/pogo/entity/system"
	systemService "admin/business/service/system"
	"errors"

	"git.xios.club/xios/gc"

	"gorm.io/gorm"
)

func init() {
	gc.RegisterBean(new(MenuServiceImpl)).Export((*systemService.MenuService)(nil))
}

type MenuServiceImpl struct {
	Db *gorm.DB `autowire:""`
}

func (this *MenuServiceImpl) SelectMenuListByUserId(userId uint) []*systemDTO.MenuInfo {
	var menuInfos []*systemDTO.MenuInfo
	this.Db.Table("menu").Where("id in (?)",
		this.Db.Table("user_role a").Select("b.menu_id").
			Joins("left join role_menu b on a.role_id = b.role_id").
			Where("a.user_id = ?", userId).Group("b.menu_id"),
	).Where("status = 0").Where("type < 2").
		Order("`parent_id` asc, `order` asc").Find(&menuInfos)
	return menuInfos
}

func (this *MenuServiceImpl) SelectMenuList() []*systemDTO.MenuInfo {
	var menuInfos []*systemDTO.MenuInfo
	this.Db.Model(systemEntity.Menu{}).Order("`parent_id` asc, `order` asc").Find(&menuInfos)
	return menuInfos
}

func (this *MenuServiceImpl) InsertMenu(menuInfo systemBO.AddMenuInfo) error {
	if !this.hasParentId(menuInfo.ParentId) {
		return errors.New("未找到父级菜单")
	}
	menu := systemEntity.Menu{
		Name:      menuInfo.Name,
		Title:     menuInfo.Title,
		Icon:      menuInfo.Icon,
		Path:      menuInfo.Path,
		Query:     menuInfo.Query,
		Redirect:  menuInfo.Redirect,
		Component: menuInfo.Component,
		Order:     menuInfo.Order,
		ParentId:  menuInfo.ParentId,
		Remark:    menuInfo.Remark,
		Perm:      menuInfo.Perm,
		Type:      menuInfo.Type,
		Visible:   menuInfo.Visible,
	}
	return this.Db.Create(&menu).Error
}

func (this *MenuServiceImpl) UpdateMenu(menuInfo systemBO.EditMenuInfo) error {
	if !this.hasParentId(menuInfo.ParentId) {
		return errors.New("未找到父级菜单")
	}
	menu := systemEntity.Menu{
		Model:     gorm.Model{ID: menuInfo.ID},
		Name:      menuInfo.Name,
		Title:     menuInfo.Title,
		Icon:      menuInfo.Icon,
		Path:      menuInfo.Path,
		Query:     menuInfo.Query,
		Redirect:  menuInfo.Redirect,
		Component: menuInfo.Component,
		Order:     menuInfo.Order,
		ParentId:  menuInfo.ParentId,
		Remark:    menuInfo.Remark,
		Perm:      menuInfo.Perm,
		Type:      menuInfo.Type,
		Visible:   menuInfo.Visible,
	}
	return this.Db.Updates(menu).Error
}

func (this *MenuServiceImpl) SelectMenuById(menuId uint) (*systemDTO.MenuInfo, error) {
	menuInfo := &systemDTO.MenuInfo{}
	if err := this.Db.Model(systemEntity.Menu{}).Find(menuInfo, menuId).Error; err != nil {
		return nil, err
	}
	return menuInfo, nil
}

func (this *MenuServiceImpl) DeleteMenuById(menuId uint) error {
	return this.Db.Delete(&systemEntity.Menu{}, menuId).Error
}

func (this *MenuServiceImpl) hasParentId(parentId uint) bool {
	if err := this.Db.First(&systemEntity.Menu{ParentId: parentId}).Error; err != nil {
		return false
	}
	return true
}
