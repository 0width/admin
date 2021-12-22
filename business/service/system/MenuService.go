package SystemService

import (
	systemBO "admin/business/pogo/bo/system"
	systemDTO "admin/business/pogo/dto/system"
)

type MenuService interface {
	SelectMenuList(userId uint) []*systemDTO.MenuInfo
	InsertMenu(info systemBO.AddMenuInfo) error
	UpdateMenu(info systemBO.EditMenuInfo) error
}
