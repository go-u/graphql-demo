package store

import (
	"context"
	"github.com/volatiletech/sqlboiler/boil"
	domain "server/domain/model"
	"server/domain/repository"
	"server/infrastructure/store/mysql/models"
)

func NewUserRepository(sqlHandler SqlHandler) repository.UserRepository {
	userRepository := UserStore{sqlHandler}
	return &userRepository
}

type UserStore struct {
	SqlHandler
}

func (s *UserStore) GetMulti(ctx context.Context, uids []string) ([]*domain.User, error) {
	users, err := models.Users(models.UserWhere.UID.IN(uids)).All(ctx, s.SqlHandler.Conn)
	if err != nil {
		return nil, err
	}
	var userEntities []*domain.User
	for _, user := range users {
		userEntity := &domain.User{
			UID:       user.UID,
			CreatedAt: int(user.CreatedAt.Unix()),
		}
		userEntities = append(userEntities, userEntity)
	}
	return userEntities, nil
}

func (s *UserStore) GetByUID(ctx context.Context, uid string) (*domain.User, error) {
	user, err := models.Users(models.UserWhere.UID.EQ(uid)).One(ctx, s.SqlHandler.Conn)
	if err != nil {
		return nil, err
	}
	userEntity := &domain.User{
		UID:       user.UID,
		CreatedAt: int(user.CreatedAt.Unix()),
	}
	return userEntity, err
}

func (s *UserStore) Create(ctx context.Context, uid string) (*domain.User, error) {
	user := models.User{
		UID: uid,
	}
	err := user.Insert(ctx, s.SqlHandler.Conn, boil.Infer())

	userEntity := &domain.User{
		UID:       uid,
		CreatedAt: int(user.CreatedAt.Unix()),
	}

	return userEntity, err
}
