package usecase

import (
	"context"
	"errors"
	domain "server/domain/model"
	"server/domain/repository"
)

type TodoUsecase interface {
	CreateTodo(context.Context, *domain.Todo) (*domain.Todo, error)
	DeleteTodo(context.Context, string, uint64) error
	GetByUID(context.Context, string) ([]*domain.Todo, error)
}

type todoUsecase struct {
	Repo repository.TodoRepository
}

func NewTodoUsecase(todoRepo repository.TodoRepository) TodoUsecase {
	todoUsecase := todoUsecase{todoRepo}
	return &todoUsecase
}

func (u *todoUsecase) CreateTodo(ctx context.Context, todo *domain.Todo) (*domain.Todo, error) {
	todo, err := u.Repo.Create(ctx, todo)
	if err != nil {
		return nil, err
	}
	return todo, nil
}

func (u *todoUsecase) DeleteTodo(ctx context.Context, uid string, id uint64) error {
	todoEntity, err := u.Repo.GetByID(ctx, id)
	if err != nil {
		return err
	}
	isOwner := todoEntity.UID == uid
	if !isOwner {
		return errors.New("not owner")
	}
	return u.Repo.Delete(ctx, id)
}

func (u *todoUsecase) GetByUID(ctx context.Context, uid string) ([]*domain.Todo, error) {
	todos, err := u.Repo.GetByUID(ctx, uid)
	if err != nil {
		return nil, err
	}
	return todos, nil
}
