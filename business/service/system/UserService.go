package SystemService

import (
	commonBO "admin/business/pogo/bo/common"
	systemDTO "admin/business/pogo/dto/system"
)

type SystemUserService interface {
	SelectUserList(page *commonBO.CommonPage) []*systemDTO.SystemUserInfoDTO
	SelectUserById(userId uint) *systemDTO.SystemUserInfoDTO
}
