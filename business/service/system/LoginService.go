package SystemService

type SystemLoginService interface {
	Login(userName, Password string) (string, error)
}
