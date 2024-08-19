package main

import (
	"context"
	"log/slog"
	"os"
	"refstor/cmd/application"

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

	app := application.New(PORT)
	log.Info("Server starting on port:" + PORT)
	err = app.Start(context.TODO())
	if err != nil {
		log.Error("Start app:", err)
	}

	/*
			router := chi.NewRouter()

			// добавляем к кождому запросу request-id
			router.Use(middleware.RequestID)

			//router.Use(logger.New(log))
			//router.Use(middleware.Recoverer)
			//router.Use(middleware.URLFormat)

			// router.Post("/url", save.New(log, storage, cfg.AliasLength))

			router.Get("/", func(w http.ResponseWriter, r *http.Request) {
				w.Write([]byte("welcome"))
			})

		log.Info("Server starting on port:" + PORT)
		http.ListenAndServe(":"+PORT, router)

	*/
}
