package SystemService

import (
	systemBO "admin/business/pogo/bo/system"
	systemDTO "admin/business/pogo/dto/system"
)

type MenuService interface {
	SelectMenuListByUserId(userId uint) []*systemDTO.MenuInfo
	SelectMenuList() []*systemDTO.MenuInfo
	InsertMenu(info systemBO.AddMenuInfo) error
	UpdateMenu(info systemBO.EditMenuInfo) error
	SelectMenuById(menuId uint) (*systemDTO.MenuInfo, error)
	DeleteMenuById(menuId uint) error
}
