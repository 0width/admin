package systemBO

type UserInfo struct {
	Name     string `json:"name" binding:"required" required_err:"用户名不能为空"`
	Password string `json:"password" binding:"required" required_err:"密码不能为空"`
	NickName string `json:"nick_name"`
	Email    string `json:"email"`
	Phone    string `json:"phone" binding:"required" required_err:"手机号不能为空"`
	Sex      int    `json:"sex"`
	Avatar   string `json:"avatar"`
	Remark   string `json:"remark"`
	DeptId   uint   `json:"dept_id"`
}

type UpdateUserInfo struct {
	ID       uint   `json:"id" binding:"required" required_err:"用ID不能为空"`
	NickName string `json:"nick_name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Sex      int    `json:"sex"`
	Avatar   string `json:"avatar"`
	Remark   string `json:"remark"`
	RoleIds  []uint `json:"roleIds"`
}
