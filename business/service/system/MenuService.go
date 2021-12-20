package SystemService

import systemDTO "admin/business/pogo/dto/system"

type SystemMenuService interface {
	SelectMenuList(userId uint) []*systemDTO.SystemMenuInfoDTO
}
