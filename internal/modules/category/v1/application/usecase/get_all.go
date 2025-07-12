package usecase

import (
	"context"
	"profotdel-rest/internal/modules/category/v1/domain"
)

type GetAllUseCase interface {
	Execute(ctx context.Context) ([]*domain.Category, error)
}

type getAllUseCase struct {
	repo domain.Repository
}

func NewGetAllUseCase(repo domain.Repository) GetAllUseCase {
	return &getAllUseCase{
		repo: repo,
	}
}

func (u *getAllUseCase) Execute(ctx context.Context) ([]*domain.Category, error) {
	models, err := u.repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return models, nil
}
