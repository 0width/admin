package systemService

type LoginService interface {
	Login(userName, Password string) (string, error)
	Logout(userId uint) error
}
