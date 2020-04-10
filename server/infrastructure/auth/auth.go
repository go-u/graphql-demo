package auth

import (
	"context"
	fbauth "firebase.google.com/go/auth"
	"log"
	"server/domain/service"
)

type AuthService interface {
	Verify(string) (uid string, err error)
}

type authService struct {
	Client fbauth.Client
}

func NewAuthService(client fbauth.Client) service.AuthService {
	authService := authService{client}
	return &authService
}

func (a *authService) Verify(token string) (string, error) {
	verified_token, err := a.Client.VerifyIDToken(context.Background(), token)
	if err != nil {
		log.Printf("error verifying ID token: %v\n", err)
		return "", err
	}
	return verified_token.UID, nil
}
