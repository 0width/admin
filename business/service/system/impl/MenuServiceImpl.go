package impl

import (
	"admin/business/pogo/dto/system"
	system2 "admin/business/service/system"

	"git.xios.club/xios/gc"

	"gorm.io/gorm"
)

func init() {
	gc.RegisterBean(new(MenuServiceImpl)).Export((*system2.MenuService)(nil))
}

type MenuServiceImpl struct {
	Db *gorm.DB `autowire:""`
}

func (this *MenuServiceImpl) SelectMenuList(userId uint) []*system.MenuInfo {
	var menuInfos []*system.MenuInfo
	this.Db.Table("menu").Where("id in (?)",
		this.Db.Table("user_role a").Select("b.menu_id").
			Joins("left join role_menu b on a.role_id = b.role_id").
			Where("a.user_id = ?", userId).Group("b.menu_id"),
	).Where("status = 0").Order("`level` asc, `order` asc").Find(&menuInfos)
	return menuInfos
}
