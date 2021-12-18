package system

import (
	"admin/business/pogo/bo/common"
	"admin/business/pogo/dto/system"
)

type UserService interface {
	SelectUserList(page *common.Page) []*system.UserInfo
	SelectUserById(userId uint) *system.UserInfo
}
