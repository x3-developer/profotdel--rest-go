package rest

import (
	"github.com/go-chi/chi/v5"
	"profotdel-rest/internal/modules/category/v1/application/usecase"
	"profotdel-rest/internal/modules/category/v1/infrastructure/persistence"
	"profotdel-rest/internal/modules/category/v1/interfaces/rest/handlers"
	"profotdel-rest/internal/shared/application/container"
)

func CategoryV1Routes(r *chi.Mux, container *container.Container) {
	repo := repository.NewRepository(container.PostgresDB)
	createUC := usecase.NewCreateUseCase(repo)
	getAllUC := usecase.NewGetAllUseCase(repo)
	h := handlers.NewHandler(createUC, getAllUC)

	r.Post("/category/create", h.Create)
	r.Get("/category", h.GetAll)
}
