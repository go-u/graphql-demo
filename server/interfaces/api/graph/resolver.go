package graph

import (
	usecase "server/application/usercase"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	AuthUsecase   usecase.AuthUsecase
	TodoUsecase   usecase.TodoUsecase
	ToggleUsecase usecase.ToggleUsecase
	UserUsecase   usecase.UserUsecase
}
