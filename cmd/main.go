package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"refstor/cmd/application"

	"github.com/lpernett/godotenv"
)

func main() {
	var log *slog.Logger
	log = slog.New(slog.NewTextHandler(os.Stderr, nil))

	// TODO - create config file
	err := godotenv.Load(".env")
	if err != nil {
		log.Error("Error loading .env file")
	}
	PORT := os.Getenv("PORT")

	config := make(map[string]string)
	config["port"] = PORT

	app := application.New(config)

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	log.Info("Server starting on port:" + PORT)
	err = app.Start(ctx)
	if err != nil {
		log.Error("Start app:", err)
	}

}
