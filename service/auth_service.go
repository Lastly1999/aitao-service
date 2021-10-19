package service

type AuthService struct {
}
type IAuthService interface {
	GetUserAuth()
}

func (authService *AuthService) GetUserAuth() {
	//TODO implement me
	panic("implement me")
}
