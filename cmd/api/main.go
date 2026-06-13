package main

import (
	"log"
	"net/http"
	"subs-app/internal/config"
	"subs-app/internal/database"
	"subs-app/internal/handlers"
	"subs-app/internal/repositories"
	"subs-app/internal/services"

	"github.com/go-chi/chi/v5"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"
)

func main() {
	cfg := config.Load()

	db, err := database.New(cfg)
	if err != nil {
		log.Fatalf("database init error: %v", err)
	}

	r := chi.NewRouter()
	r.Use(chiMiddleware.RequestID)
	r.Use(chiMiddleware.Recoverer)
	// r.Use(appMiddleware.Logger)

	repo := repositories.NewRepository(db)
	service := services.NewService(repo)
	handler := handlers.NewHandler(service)

	r.Post("/create_subscription", handler.CreateSub)
	r.Get("/subscriptions", handler.GetSub)
	r.Put("/subscriptions", handler.UpdateSub)
	r.Delete("/subscriptions", handler.DeleteSub)
	r.Get("/subscriptions/aggregate", handler.AggregateSubs)

	addr := ":" + cfg.AppPort
	log.Printf("server started on %v", addr)
	if err = http.ListenAndServe(addr, r); err != nil {
		log.Fatalf("server stopped: %v", err)
	}
}
