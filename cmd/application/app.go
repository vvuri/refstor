package application

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"net/http"
	"time"
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

	// DB Redis
	err := a.rdb.Ping(ctx).Err()
	if err != nil {
		return fmt.Errorf("Failed connect ot Redis: %w", err)
	}

	defer func() {
		if err := a.rdb.Close(); err != nil {
			fmt.Errorf("Failed to close Radis: %w", err)
		}
	}()

	ch := make(chan error, 1)

	// run api server
	go func() {
		err = server.ListenAndServe()
		if err != nil {
			ch <- fmt.Errorf("Failed to start server: %w", err)
		}
		close(ch)
	}()

	// stop service
	select {
	case err = <-ch:
		return err
	case <-ctx.Done():
		timeout, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()

		return server.Shutdown(timeout)
	}

	return nil
}
