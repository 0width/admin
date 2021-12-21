package SystemService

import (
	commonBO "admin/business/pogo/bo/common"
	systemDTO "admin/business/pogo/dto/system"
)

type UserService interface {
	SelectUserList(page *commonBO.Page) []*systemDTO.UserInfo
	SelectUserById(userId uint) *systemDTO.UserInfo
}
