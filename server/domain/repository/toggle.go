package repository

import (
	"context"
	"server/domain/model"
)

type ToggleRepository interface {
	GetAll(ctx context.Context) ([]*model.Toggle, error)
	GetByUID(ctx context.Context, uid string) (*model.Toggle, error)
	Switch(ctx context.Context, uid string, enable bool) (*model.Toggle, error)
}
