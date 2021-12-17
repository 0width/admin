package system

type LoginService interface {
	Login(userName, Password string) (string, error)
}
