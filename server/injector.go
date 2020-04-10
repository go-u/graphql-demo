package main

import (
	"server/application/usercase"
	"server/infrastructure/auth"
	"server/infrastructure/store/mysql"
	"server/interfaces/api"
)

//func getProjectId() string {
//	PROJECT_ID := os.Getenv("GOOGLE_CLOUD_PROJECT")
//	if PROJECT_ID == "" {
//		log.Fatalln("Failed to Get PROJECT_ID 'GOOGLE_CLOUD_PROJECT'\n If this is local test, Set 'appname-local' as GOOGLE_CLOUD_PROJECT")
//	}
//	log.Println("PROJECT_ID: ", PROJECT_ID)
//	return PROJECT_ID
//}

func getApiConfig() api.Config {
	projectID := "sample"

	// infra
	sqlHandler := store.NewSqlHandler(projectID)
	authCient := auth.NewClient(projectID)

	// repository & service
	authService := auth.NewAuthService(*authCient)
	todoRepository := store.NewTodoRepository(*sqlHandler)
	toggleRepository := store.NewToggleRepository(*sqlHandler)
	userRepository := store.NewUserRepository(*sqlHandler)

	// usecase
	authUsecase := usecase.NewAuthUsecase(authService)
	todoUsecase := usecase.NewTodoUsecase(todoRepository)
	toggleUsecase := usecase.NewToggleUsecase(toggleRepository)
	userUsecase := usecase.NewUserUsecase(userRepository)

	config := api.Config{
		AuthUsecase:   authUsecase,
		TodoUsecase:   todoUsecase,
		ToggleUsecase: toggleUsecase,
		UserUsecase:   userUsecase,
	}
	return config
}
