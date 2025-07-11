package router

import (
	"github.com/go-chi/chi/v5"
	"net/http"
	"profotdel-rest/config"
	"profotdel-rest/internal/container"
	"profotdel-rest/internal/middleware"
	"profotdel-rest/internal/modules/category/v1/delivery"
)

func NewRouter(cfg *config.Config, container *container.Container) *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	delivery.CategoryV1Routes(r, container)

	r.Handle("/public/*", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))

	return r
}
