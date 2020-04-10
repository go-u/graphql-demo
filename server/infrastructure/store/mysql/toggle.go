package store

import (
	"context"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"
	domain "server/domain/model"
	"server/domain/repository"
	"server/infrastructure/store/mysql/models"
)

func NewToggleRepository(sqlHandler SqlHandler) repository.ToggleRepository {
	toggleRepository := ToggleStore{sqlHandler}
	return &toggleRepository
}

type ToggleStore struct {
	SqlHandler
}

func (s *ToggleStore) GetAll(ctx context.Context) ([]*domain.Toggle, error) {
	toggles, err := models.Toggles(qm.Limit(100)).All(ctx, s.SqlHandler.Conn)
	if err != nil {
		return nil, err
	}
	var toggleEntities []*domain.Toggle
	for _, toggle := range toggles {
		toggleEntity := &domain.Toggle{
			UID:    toggle.UID,
			Enable: toggle.Enable,
		}
		toggleEntities = append(toggleEntities, toggleEntity)

	}
	return toggleEntities, nil
}

func (s *ToggleStore) GetByUID(ctx context.Context, uid string) (*domain.Toggle, error) {
	toggle, err := models.Toggles(models.ToggleWhere.UID.EQ(uid)).One(ctx, s.SqlHandler.Conn)
	if err != nil {
		return nil, err
	}
	toggleEntity := &domain.Toggle{
		UID:    toggle.UID,
		Enable: toggle.Enable,
	}
	return toggleEntity, err
}

func (s *ToggleStore) Switch(ctx context.Context, uid string, enable bool) (*domain.Toggle, error) {
	toggle := models.Toggle{
		UID:    uid,
		Enable: enable,
	}
	err := toggle.Upsert(ctx, s.SqlHandler.Conn, boil.Blacklist(), boil.Blacklist())

	toggleEntity := &domain.Toggle{
		UID:    uid,
		Enable: enable,
	}

	return toggleEntity, err
}
