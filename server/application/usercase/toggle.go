package usecase

import (
	"context"
	domain "server/domain/model"
	"server/domain/repository"
)

type ToggleUsecase interface {
	GetAll(context.Context) ([]*domain.Toggle, error)
	GetByUID(context.Context, string) (*domain.Toggle, error)
	Switch(context.Context, string, bool) (*domain.Toggle, error)
}

type toggleUsecase struct {
	Repo repository.ToggleRepository
}

func NewToggleUsecase(toggleRepo repository.ToggleRepository) ToggleUsecase {
	toggleUsecase := toggleUsecase{toggleRepo}
	return &toggleUsecase
}

func (u *toggleUsecase) GetAll(ctx context.Context) ([]*domain.Toggle, error) {
	return u.Repo.GetAll(ctx)
}

func (u *toggleUsecase) GetByUID(ctx context.Context, uid string) (*domain.Toggle, error) {
	toggle, err := u.Repo.GetByUID(ctx, uid)
	if err != nil {
		return nil, err
	}
	return toggle, nil
}

func (u *toggleUsecase) Switch(ctx context.Context, uid string, enable bool) (*domain.Toggle, error) {
	toggle, err := u.Repo.Switch(ctx, uid, enable)
	if err != nil {
		return nil, err
	}
	return toggle, nil
}
