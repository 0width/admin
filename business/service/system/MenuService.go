package SystemService

import systemDTO "admin/business/pogo/dto/system"

type MenuService interface {
	SelectMenuList(userId uint) []*systemDTO.MenuInfo
}
