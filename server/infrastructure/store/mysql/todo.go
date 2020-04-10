package store

import (
	"context"
	"github.com/volatiletech/sqlboiler/boil"
	domain "server/domain/model"
	"server/domain/repository"
	"server/infrastructure/store/mysql/models"
)

func NewTodoRepository(sqlHandler SqlHandler) repository.TodoRepository {
	todoRepository := TodoStore{sqlHandler}
	return &todoRepository
}

type TodoStore struct {
	SqlHandler
}

func (s *TodoStore) GetByID(ctx context.Context, id uint64) (*domain.Todo, error) {
	todo, err := models.FindTodo(ctx, s.SqlHandler.Conn, id)
	if err != nil {
		return nil, err
	}
	todoEntity := &domain.Todo{
		ID:   todo.ID,
		UID:  todo.UID,
		Text: todo.Text,
	}
	return todoEntity, err
}

func (s *TodoStore) GetByUID(ctx context.Context, uid string) ([]*domain.Todo, error) {
	todos, err := models.Todos(models.TodoWhere.UID.EQ(uid)).All(ctx, s.SqlHandler.Conn)
	if err != nil {
		return nil, err
	}

	var todoEntities = []*domain.Todo{}
	for _, todo := range todos {
		todoEntity := &domain.Todo{
			ID:   todo.ID,
			UID:  todo.UID,
			Text: todo.Text,
		}
		todoEntities = append(todoEntities, todoEntity)
	}
	return todoEntities, err
}

func (s *TodoStore) Create(ctx context.Context, todoEntity *domain.Todo) (*domain.Todo, error) {
	todo := &models.Todo{
		ID:   todoEntity.ID,
		UID:  todoEntity.UID,
		Text: todoEntity.Text,
	}
	err := todo.Insert(ctx, s.SqlHandler.Conn, boil.Infer())
	// add id
	todoEntity.ID = todo.ID
	return todoEntity, err
}

func (s *TodoStore) Delete(ctx context.Context, id uint64) error {
	todo := &models.Todo{
		ID: id,
	}
	_, err := todo.Delete(ctx, s.SqlHandler.Conn)
	return err
}
