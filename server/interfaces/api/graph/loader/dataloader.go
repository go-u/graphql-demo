package loader

import (
	"context"
	"github.com/graph-gophers/dataloader"
	"log"
	usecase "server/application/usercase"
)

const (
	UsersLoaderKey = "UsersLoaderKey"
)

func GetUsersBatchFunc(userUsecase usecase.UserUsecase) dataloader.BatchFunc {
	return func(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
		var results []*dataloader.Result
		var result *dataloader.Result

		var IDs []string
		for _, key := range keys {
			ID := key.String()
			IDs = append(IDs, ID)
		}

		users, err := userUsecase.GetMulti(ctx, IDs)
		if err != nil {
			log.Println(ctx, err.Error())
			result = handleError(ctx, err)
			results = append(results, result)
		}

		for _, user := range users {
			user.BatchSize = len(users)
			result := dataloader.Result{
				Data:  user,
				Error: nil,
			}
			results = append(results, &result)
		}

		log.Printf("[GetUsersBatchFunc] batch size: %d", len(results))
		return results
	}
}

func handleError(ctx context.Context, err error) *dataloader.Result {
	log.Println(ctx, err.Error())
	var result dataloader.Result
	result.Error = err
	return &result
}
