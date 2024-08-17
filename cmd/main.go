package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log/slog"
	"net/http"
	"os"

	"github.com/lpernett/godotenv"
)

func main() {
	var log *slog.Logger
	log = slog.New(slog.NewTextHandler(os.Stderr, nil))

	err := godotenv.Load(".env")
	if err != nil {
		log.Error("Error loading .env file")
	}
	PORT := os.Getenv("PORT")

	router := chi.NewRouter()

	// добавляем к кождому запросу request-id
	//router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	//router.Use(logger.New(log))
	//router.Use(middleware.Recoverer)
	//router.Use(middleware.URLFormat)

	// router.Post("/url", save.New(log, storage, cfg.AliasLength))

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})

	log.Info("Server starting on port:" + PORT)
	http.ListenAndServe(":"+PORT, router)
}
