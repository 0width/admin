package systemEntity

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `gorm:"size:255;index"`
	Password string `gorm:"size:64" json:"-"`
	NickName string `gorm:"size:255;comment:昵称"`
	Email    string `gorm:"size:255"`
	Phone    string `gorm:"size:16"`
	Sex      int    `gorm:"type:tinyint;comment:1: 男 2: 女 3: 未知"`
	Status   int    `gorm:"type:tinyint;default:0;comment:0: 正常 1: 停用" json:"-"`
	Avatar   string `gorm:"size:255"`
	Remark   string `gorm:"size:1023;comment:备注"`
	DeptId   uint   `gorm:"comment:部门ID"`
	Dept     *Dept
	Roles    []*Role `gorm:"many2many:user_role"`
}

func (User) TableName() string {
	return "user"
}
