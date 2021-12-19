package system

import "admin/business/pogo/dto/system"

type MenuService interface {
	SelectMenuList(userId uint) []*system.MenuInfo
}
