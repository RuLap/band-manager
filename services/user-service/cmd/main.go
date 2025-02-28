package main

import (
	"band-manager/services/user-service/internal/config"
	"band-manager/services/user-service/internal/handlers"
	"band-manager/services/user-service/internal/repository"
	"band-manager/services/user-service/internal/services"
	"band-manager/services/user-service/internal/storage/postgres"
	"band-manager/services/user-service/pkg/jwt_helper"
	"band-manager/services/user-service/pkg/recovery"
	"github.com/go-chi/chi/v5"
	"log/slog"
	"net/http"
)

func main() {
	cfg := config.MustLoad()
	slog.Info("loaded config successfuly")

	postgres.InitDB(cfg.PostgresConnString)
	slog.Info("init postgres connection successfully")

	if err := jwt_helper.NewJwtHelper(cfg.JWT.Secret); err != nil {
		slog.Error("failed to init JWT", "error", err)
		return
	}
	slog.Info("generated JWT successfuly")

	router := chi.NewRouter()
	slog.Info("init chi router successfuly")

	http.HandleFunc("/panic", recovery.Middleware(panicHandler))

	userRepo := repository.NewUserRepository(postgres.GetDB())
	slog.Info("init repositories successfuly")

	userService := services.NewUserService(userRepo)
	slog.Info("init services successfuly")

	userHandler := handlers.NewUserHandler(userService)
	slog.Info("init handlers successfuly")

	router.Group(func(r chi.Router) {
		r.Post("/login", userHandler.Login)
		r.Post("/register", userHandler.Register)
		r.Post("/users/{id}", userHandler.GetUser)
	})

	slog.Info("init routes successfuly")

	server := &http.Server{
		Addr:         cfg.HTTPServer.Address,
		Handler:      router,
		ReadTimeout:  cfg.HTTPServer.Timeout,
		WriteTimeout: cfg.HTTPServer.Timeout,
		IdleTimeout:  cfg.HTTPServer.IdleTimeout,
	}

	if err := server.ListenAndServe(); err != nil {
		slog.Info("server error", "error", err)
	}
}

func panicHandler(w http.ResponseWriter, r *http.Request) {
	panic("Something went wrong!")
}
