package delivery

import (
	"github.com/go-chi/chi/v5"
	"profotdel-rest/internal/container"
	"profotdel-rest/internal/modules/category/v1/repository"
	"profotdel-rest/internal/modules/category/v1/usecase"
)

func CategoryV1Routes(r *chi.Mux, container *container.Container) {
	repo := repository.NewRepository(container.DB)
	createUC := usecase.NewCreateUseCase(repo)
	getAllUC := usecase.NewGetAllUseCase(repo)
	h := NewHandler(createUC, getAllUC)

	r.Post("/category/create", h.Create)
	r.Get("/category", h.GetAll)
}
