package SystemService

import (
	commonBO "admin/business/pogo/bo/common"
	systemBO "admin/business/pogo/bo/system"
	systemDTO "admin/business/pogo/dto/system"
)

type UserService interface {
	SelectUserList(page *commonBO.Page) []*systemDTO.UserInfo
	SelectUserById(userId uint) *systemDTO.UserInfo
	UpdateRoles(userId uint, roleId ...uint) error
	UpdateUser(info systemBO.UpdateUserInfo) error
}
