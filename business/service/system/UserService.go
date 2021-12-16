package system

import "admin/business/pogo/entity"

type UserService interface {
	SelectUserList() []entity.User
}
