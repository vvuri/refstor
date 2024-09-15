package application

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
	"refstor/cmd/handler"
	"refstor/cmd/repository"
)

func (a *App) loadRoutes() {
	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})

	router.Route("/api/img", a.loadImageRoutes)

	a.router = router
}

func (a *App) loadImageRoutes(router chi.Router) {
	imageHandler := &handler.Image{
		Repo: &repository.RedisRepo{
			Client: a.rdb,
		},
	}

	router.Post("/", imageHandler.Create)
	router.Get("/", imageHandler.List)
	router.Get("/{id}", imageHandler.ImageByID)
	router.Put("/{id}", imageHandler.UpdateByID)
	router.Delete("/{id}", imageHandler.Delete)
}
