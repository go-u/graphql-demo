package middleware

import (
	"context"
	"github.com/graph-gophers/dataloader"
	"net/http"
	"server/interfaces/api"
	"server/interfaces/api/graph/loader"
)

func AddDataloaderContext(config api.Config) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			// requestの度に、ユーザー個別の一貫したDataloader用コンテキストを生成
			newLoader := dataloader.NewBatchedLoader(loader.GetUsersBatchFunc(config.UserUsecase))

			ctx := context.WithValue(
				r.Context(),
				loader.UsersLoaderKey,
				newLoader)
			next.ServeHTTP(w, r.WithContext(ctx))
		}
		return http.HandlerFunc(fn)
	}
}
