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
}
