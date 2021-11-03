package di

import (
	"context"
	"encoding/json"
	"net/http"

	"user-service/internal/api"
	"user-service/internal/postgres"
	"user-service/internal/postgres/database"

	"github.com/jackc/pgx/v4/pgxpool"
)

func NewRouter(connection *pgxpool.Pool, version, env string) http.Handler {
	db := database.New(connection)

	userRepository := postgres.NewUserRepository(db)
	userApiService := api.NewUserApiService(userRepository)
	userApiController := api.NewUserApiController(userApiService)

	router := api.NewRouter(userApiController)

	router.HandleFunc("/health", func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("content-type", "application/json")
		writer.WriteHeader(http.StatusOK)
		writer.Write([]byte(`{"status":"ok"}`))
	})

	router.HandleFunc("/ready", func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("content-type", "application/json")
		err := connection.Ping(context.Background())
		if err == nil {
			writer.WriteHeader(http.StatusOK)
			writer.Write([]byte(`{"status":"ok"}`))
		} else {
			writer.WriteHeader(http.StatusServiceUnavailable)
			writer.Write([]byte(`{"status":"not available"}`))
		}
	})

	router.HandleFunc("/version", func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("content-type", "application/json")
		writer.WriteHeader(http.StatusOK)
		response, _ := json.Marshal(struct {
			Environment string `json:"environment"`
			Version     string `json:"version"`
		}{
			Environment: env,
			Version:     version,
		})
		writer.Write(response)
	})

	return router
}
