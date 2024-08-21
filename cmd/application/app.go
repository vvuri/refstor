package application

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"net/http"
)

type App struct {
	router http.Handler
	port   string
	rdb    *redis.Client
}

func New(config map[string]string) *App {
	app := &App{
		router: loadRoutes(),
		port:   config["port"],
		rdb:    redis.NewClient(&redis.Options{}),
	}
	return app
}

func (a *App) Start(ctx context.Context) error {
	server := &http.Server{
		Addr:    ":" + a.port,
		Handler: a.router,
	}

	err := a.rdb.Ping(ctx).Err()
	if err != nil {
		return fmt.Errorf("Failed connect ot Redis: %w", err)
	}

	err = server.ListenAndServe()
	if err != nil {
		return fmt.Errorf("Failed to start server: %w", err)
	}
	return nil
}
