package router

import (
	"github.com/go-chi/chi/v5"
	"net/http"
	"profotdel-rest/config"
	"profotdel-rest/internal/container"
)

func NewRouter(cfg *config.Config, container *container.Container) *chi.Mux {
	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("Welcome to the Profotdel API!"))
		if err != nil {
			return
		}
	})

	return r
}
