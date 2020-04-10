package main

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/volatiletech/sqlboiler/boil"
	"log"
	"net/http"
	"os"
	"server/interfaces/api"
	middlewareCustom "server/interfaces/api/middleware"
)

func init() {
	boil.DebugMode = false
}

func main() {
	// Dependency Injection
	handlerConfig := getApiConfig()
	apiHandler := api.NewHandler(handlerConfig)

	// router
	r := chi.NewRouter()
	r.Use(cors.New(
		cors.Options{
			AllowedOrigins: []string{"*"}, // for development
			AllowedMethods: []string{"POST"},
			AllowedHeaders: []string{"Accept", "Content-Type", "JWT"},
			MaxAge:         3600,
		}).Handler)
	r.Use(middleware.Logger)
	r.Use(middlewareCustom.LimitSizeByContentLengthHeader)
	r.Use(middlewareCustom.GetCloudFlareIp)
	r.Use(middlewareCustom.AddJwtContext)
	r.Use(middlewareCustom.AddDataloaderContext(handlerConfig)) // dataloader
	r.Use(middleware.Recoverer)
	r.Mount("/api", apiHandler)

	// start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}
	log.Printf("Listening on port %s", port)
	err := http.ListenAndServe(":"+port, r)
	if err != nil {
		log.Fatalln(err)
	}
}
