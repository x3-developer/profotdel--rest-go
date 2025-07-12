package router

import (
	"github.com/go-chi/chi/v5"
	"net/http"
	"profotdel-rest/config"
	"profotdel-rest/internal/modules/category/v1/interfaces/rest"
	"profotdel-rest/internal/shared/application/container"
	"profotdel-rest/internal/shared/application/middleware"
)

func NewRouter(cfg *config.Config, container *container.Container) *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.CORSMiddleware(cfg.CORS))
	r.Use(middleware.APIMiddleware(cfg.AuthAppKey))

	rest.CategoryV1Routes(r, container)

	r.Handle("/public/*", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))

	return r
}
