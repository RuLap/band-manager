package main

import (
	"band-manager/services/user-service/internal/config"
	"band-manager/services/user-service/pkg/jwt_helper"
	"band-manager/services/user-service/pkg/recovery"
	"github.com/go-chi/chi/v5"
	"log/slog"
	"net/http"
)

func main() {
	cfg := config.MustLoad()
	slog.Info("loaded config successfuly")

	if err := jwt_helper.NewJwtHelper(cfg.JWT.Secret); err != nil {
		slog.Error("failed to init JWT", "error", err)
		return
	}
	slog.Info("generated JWT successfuly")

	router := chi.NewRouter()
	slog.Info("init chi router successfuly")

	http.HandleFunc("/panic", recovery.Middleware(panicHandler))

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
