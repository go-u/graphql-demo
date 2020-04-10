package repository

import (
	"context"
	"server/domain/model"
)

type TodoRepository interface {
	GetByID(ctx context.Context, id uint64) (*model.Todo, error)
	GetByUID(ctx context.Context, uid string) ([]*model.Todo, error)
	Create(ctx context.Context, todo *model.Todo) (*model.Todo, error)
	Delete(ctx context.Context, id uint64) error
}
