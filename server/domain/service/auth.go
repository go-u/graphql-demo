package service

type AuthService interface {
	Verify(jwt string) (uid string, err error)
}
