package graph

import (
	"context"
	"errors"
	domain "server/domain/model"
)

// chi middlewareで事前にjwtをcontextに格納
func (r *Resolver) getUIDFromContext(ctx context.Context) (string, error) {
	jwt, ok := ctx.Value("jwt").(string)
	if !ok {
		return "", errors.New("no jwt in context")
	}
	uid, err := r.AuthUsecase.Verify(jwt)
	return uid, err
}

func (r *Resolver) getUserFromContext(ctx context.Context) (*domain.User, error) {
	uid, err := r.getUIDFromContext(ctx)
	if err != nil {
		return nil, err
	}
	return r.UserUsecase.GetByUID(ctx, uid)
}
