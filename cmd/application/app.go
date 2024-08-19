package application

import (
	"context"
	"fmt"
	"net/http"
)

type App struct {
	router http.Handler
	port   string
}

func New(port string) *App {
	app := &App{
		router: loadRoutes(),
		port:   port,
	}
	return app
}

func (a *App) Start(ctx context.Context) error {
	server := &http.Server{
		Addr:    ":" + a.port,
		Handler: a.router,
	}
	err := server.ListenAndServe()
	if err != nil {
		return fmt.Errorf("Failed to start server: %w", err)
	}
	return nil
}
