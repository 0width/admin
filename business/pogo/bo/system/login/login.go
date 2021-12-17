package login

type Login struct {
	UserName string `json:"username" binding:"required" required_err:"用户名不能为空"`
	Password string `json:"password" binding:"required" required_err:"密码不能为空"`
}
