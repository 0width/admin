package systemService

type LoginService interface {
	Login(userName, Password string) (string, error)
}
