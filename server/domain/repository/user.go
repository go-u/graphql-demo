package repository

import (
	"context"
	"server/domain/model"
)

type UserRepository interface {
	GetMulti(ctx context.Context, uids []string) ([]*model.User, error)
	GetByUID(ctx context.Context, uid string) (*model.User, error)
	Create(ctx context.Context, uid string) (*model.User, error)
}
