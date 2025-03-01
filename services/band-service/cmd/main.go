package main

import (
	"band-manager/services/band-service/internal/config"
	"band-manager/services/band-service/internal/handlers"
	"band-manager/services/band-service/internal/repository"
	"band-manager/services/band-service/internal/services"
	"band-manager/services/band-service/internal/storage/postgres"
	"github.com/RuLap/band-manager/pkg/auth"
	"github.com/RuLap/band-manager/pkg/jwt_helper"
	"github.com/RuLap/band-manager/pkg/recovery"
	"github.com/go-chi/chi/v5"
	"log/slog"
	"net/http"
)

func main() {
	cfg := config.MustLoad()
	slog.Info("loaded config successfuly")

	storage, err := postgres.InitDB(cfg.PostgresConnString)
	if err != nil {
		slog.Error("failed to init postgres", "error", err)
	}
	slog.Info("init postgres connection successfully")

	router := chi.NewRouter()
	slog.Info("init chi router successfuly")

	http.HandleFunc("/panic", recovery.Middleware(panicHandler))

	bandRepo := repository.NewBandRepository(storage.Database())
	memberRepo := repository.NewMemberRepository(storage.Database())

	bandService := services.NewBandService(bandRepo, memberRepo)
	memberService := services.NewMemberService(memberRepo)

	bandHandler := handlers.NewBandHandler(bandService)
	memberHandler := handlers.NewMemberHandler(memberService)

	router.Group(func(r chi.Router) {
		r.Use(
			jwt_helper.Middleware,
			auth.Middleware,
		)

		r.Post("/bands", bandHandler.Create)
		r.Get("/bands/id", bandHandler.GetBand)

		r.Get("/members", memberHandler.GetMember)
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
