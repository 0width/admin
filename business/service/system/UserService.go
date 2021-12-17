package system

import "admin/business/pogo/entity"

type UserService interface {
	SelectUserList() []*entity.User
	SelectUserById(userId uint) *entity.User
}
