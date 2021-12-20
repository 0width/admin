package SystemServiceImpl

import (
	systemDTO "admin/business/pogo/dto/system"
	SystemService "admin/business/service/system"

	"git.xios.club/xios/gc"

	"gorm.io/gorm"
)

func init() {
	gc.RegisterBean(new(SystemMenuServiceImpl)).Export((*SystemService.SystemMenuService)(nil))
}

type SystemMenuServiceImpl struct {
	Db *gorm.DB `autowire:""`
}

func (this *SystemMenuServiceImpl) SelectMenuList(userId uint) []*systemDTO.SystemMenuInfoDTO {
	var menuInfos []*systemDTO.SystemMenuInfoDTO
	this.Db.Table("menu").Where("id in (?)",
		this.Db.Table("user_role a").Select("b.menu_id").
			Joins("left join role_menu b on a.role_id = b.role_id").
			Where("a.user_id = ?", userId).Group("b.menu_id"),
	).Where("status = 0").Order("`level` asc, `order` asc").Find(&menuInfos)
	return menuInfos
}
