package SystemService

import (
	systemBO "admin/business/pogo/bo/system"
	systemDTO "admin/business/pogo/dto/system"
)

type RoleService interface {
	SelectRoleList(userId uint) ([]systemDTO.RoleInfo, error)
	InsertRole(info systemBO.RoleInfo) error
	UpdateRole(info systemBO.RoleInfo) error
}
