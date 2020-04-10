package usecase

import (
	"context"
	domain "server/domain/model"
	"server/domain/repository"
)

type UserUsecase interface {
	Create(context.Context, string) (*domain.User, error)
	GetMulti(context.Context, []string) ([]*domain.User, error)
	GetByUID(context.Context, string) (*domain.User, error)
}

type userUsecase struct {
	Repo repository.UserRepository
}

func NewUserUsecase(userRepo repository.UserRepository) UserUsecase {
	userUsecase := userUsecase{userRepo}
	return &userUsecase
}

func (u *userUsecase) Create(ctx context.Context, uid string) (*domain.User, error) {
	user, err := u.Repo.Create(ctx, uid)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *userUsecase) GetMulti(ctx context.Context, uids []string) ([]*domain.User, error) {
	return u.Repo.GetMulti(ctx, uids)
}

func (u *userUsecase) GetByUID(ctx context.Context, uid string) (*domain.User, error) {
	user, err := u.Repo.GetByUID(ctx, uid)
	if err != nil {
		return nil, err
	}
	return user, nil
}
