package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"database/sql"
	"errors"
	domain "server/domain/model"
	"server/interfaces/api/graph/generated"
	"server/interfaces/api/graph/loader"
	"server/interfaces/api/graph/model"
	"strconv"
	"time"

	"github.com/graph-gophers/dataloader"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	// 遅延を再現
	time.Sleep(3 * time.Second)

	uid, err := r.getUIDFromContext(ctx)
	if err != nil {
		return nil, err
	}

	user, err := r.getUserFromContext(ctx)
	if user == nil || err == sql.ErrNoRows {
		user, err = r.UserUsecase.Create(ctx, uid)
	}

	if err != nil {
		return nil, err
	}

	todoEntity := &domain.Todo{
		UID:  uid,
		Text: input.Text,
	}

	todoEntity, err = r.TodoUsecase.CreateTodo(ctx, todoEntity)
	if err != nil {
		return nil, err
	}
	todo := &model.Todo{
		ID:   strconv.Itoa(int(todoEntity.ID)),
		Text: todoEntity.Text,
		UID:  todoEntity.UID,
	}

	return todo, nil
}

func (r *mutationResolver) DeleteTodo(ctx context.Context, input model.DeleteTodo) (*model.Todo, error) {
	// 遅延を再現
	time.Sleep(3 * time.Second)

	uid, err := r.getUIDFromContext(ctx)
	if err != nil {
		return nil, err
	}
	err = r.TodoUsecase.DeleteTodo(ctx, uid, uint64(input.ID))
	if err != nil {
		return nil, err
	}

	deletedTodo := &model.Todo{
		ID:   strconv.Itoa(input.ID),
		Text: "deleted",
		UID:  uid,
	}

	return deletedTodo, nil
}

func (r *mutationResolver) SwitchToggle(ctx context.Context, input model.SwitchToggle) (*model.Toggle, error) {
	// 遅延を再現
	time.Sleep(3 * time.Second)

	uid, err := r.getUIDFromContext(ctx)
	if err != nil {
		return nil, err
	}

	user, err := r.getUserFromContext(ctx)
	if user == nil || err == sql.ErrNoRows {
		user, err = r.UserUsecase.Create(ctx, uid)
	}

	toggleEntity, err := r.ToggleUsecase.Switch(ctx, uid, input.Enable)
	if err != nil {
		return nil, err
	}

	toggle := &model.Toggle{
		ID:     toggleEntity.UID,
		Enable: toggleEntity.Enable,
	}

	return toggle, nil
}

func (r *mutationResolver) SwitchToggleFail(ctx context.Context, input model.SwitchToggle) (*model.Toggle, error) {
	// 遅延を再現
	time.Sleep(3 * time.Second)

	uid, err := r.getUIDFromContext(ctx)
	if err != nil {
		return nil, err
	}

	user, err := r.getUserFromContext(ctx)
	if user == nil || err == sql.ErrNoRows {
		user, err = r.UserUsecase.Create(ctx, uid)
	}

	toggleEntity, err := r.ToggleUsecase.Switch(ctx, uid, input.Enable)
	if err != nil {
		return nil, err
	}

	toggle := &model.Toggle{
		ID:     toggleEntity.UID,
		Enable: !input.Enable, // always fail ui
	}

	return toggle, nil
}

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	uid, err := r.getUIDFromContext(ctx)
	if err != nil || len(uid) == 0 {
		return nil, err
	}

	todoEntities, err := r.TodoUsecase.GetByUID(ctx, uid)
	if err != nil {
		return nil, err
	}

	var todos = []*model.Todo{}
	for _, entity := range todoEntities {
		todo := &model.Todo{
			ID:   strconv.Itoa(int(entity.ID)),
			Text: entity.Text,
			UID:  entity.UID,
		}
		todos = append(todos, todo)
	}

	return todos, nil
}

func (r *queryResolver) Toggles(ctx context.Context) ([]*model.Toggle, error) {
	toggleEntities, err := r.ToggleUsecase.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	var toggles []*model.Toggle
	for _, toggleEntity := range toggleEntities {
		toggle := &model.Toggle{
			ID:     toggleEntity.UID,
			Enable: toggleEntity.Enable,
		}
		toggles = append(toggles, toggle)
	}
	return toggles, nil
}

func (r *queryResolver) Toggle(ctx context.Context) (*model.Toggle, error) {
	uid, err := r.getUIDFromContext(ctx)
	if err != nil || len(uid) == 0 {
		return nil, err
	}

	toggleEntity, err := r.ToggleUsecase.GetByUID(ctx, uid)
	if err == sql.ErrNoRows {
		toggle := &model.Toggle{
			ID:     uid,
			Enable: false,
		}
		return toggle, nil
	}
	if err != nil {
		return nil, err
	}

	toggle := &model.Toggle{
		ID:     toggleEntity.UID,
		Enable: toggleEntity.Enable,
	}

	return toggle, nil
}

func (r *todoResolver) User(ctx context.Context, obj *model.Todo) (*model.User, error) {
	userEntity, err := r.UserUsecase.GetByUID(ctx, obj.UID)
	if err != nil {
		return nil, err
	}

	user := &model.User{
		UID:       userEntity.UID,
		CreatedAt: userEntity.CreatedAt,
	}
	return user, nil
}

func (r *toggleResolver) User(ctx context.Context, obj *model.Toggle) (*model.User, error) {
	// context中のloaderの読込
	loaderInCtx, ok := ctx.Value(loader.UsersLoaderKey).(*dataloader.Loader)
	if !ok {
		return nil, errors.New("loader is empty")
	}

	// loaderを使ってバッチ処理を待ち受け(バッチ処理が終わるまでsleepのように待機する)
	key := obj.ID
	thunk := loaderInCtx.Load(ctx, dataloader.StringKey(key))
	// バッチ処理の結果から上記keyに受取る単一の結果を受取る
	result, err := thunk()
	if err != nil {
		return nil, err
	}

	// resultを整形してレスポンス
	userEntity, ok := result.(*domain.User)
	if !ok {
		return nil, errors.New("fail to cast user")
	}

	user := &model.User{
		UID:       userEntity.UID,
		CreatedAt: userEntity.CreatedAt,
		BatchSize: &userEntity.BatchSize,
	}
	return user, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Todo returns generated.TodoResolver implementation.
func (r *Resolver) Todo() generated.TodoResolver { return &todoResolver{r} }

// Toggle returns generated.ToggleResolver implementation.
func (r *Resolver) Toggle() generated.ToggleResolver { return &toggleResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type todoResolver struct{ *Resolver }
type toggleResolver struct{ *Resolver }
