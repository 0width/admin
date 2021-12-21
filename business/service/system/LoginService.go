package SystemService

type LoginService interface {
	Login(userName, Password string) (string, error)
}
