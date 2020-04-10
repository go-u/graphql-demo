package usecase

import (
	"server/domain/service"
)

type AuthUsecase interface {
	Verify(jwt string) (uid string, err error)
}

type authUsecase struct {
	Service service.AuthService
}

func NewAuthUsecase(authService service.AuthService) AuthUsecase {
	authUsecase := authUsecase{authService}
	return &authUsecase
}

func (u *authUsecase) Verify(jwt string) (string, error) {
	uid, err := u.Service.Verify(jwt)
	if err != nil {
		return "", err
	}
	return uid, nil
}
