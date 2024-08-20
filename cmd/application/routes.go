package application

import (
	"net/http"
	"refstor/cmd/handler"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func loadRoutes() *chi.Mux {
	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})

	router.Route("/api/img", loadImageRoutes)

	return router
}

func loadImageRoutes(router chi.Router) {
	imageHandler := &handler.Image{}

	router.Post("/", imageHandler.Create)
	router.Get("/", imageHandler.List)
	router.Get("/{id}", imageHandler.ImageByID)
	router.Put("/{id}", imageHandler.UpdateByID)
	router.Delete("/{id}", imageHandler.Delete)
}
