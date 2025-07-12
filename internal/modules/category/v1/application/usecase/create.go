package usecase

import (
	"context"
	"profotdel-rest/internal/modules/category/v1/domain"
	"profotdel-rest/internal/shared/lib/response"
)

type CreateUseCase interface {
	Execute(ctx context.Context, model *domain.Category) (*domain.Category, []response.ErrorField, error)
}

type createUseCase struct {
	repo domain.Repository
}

func NewCreateUseCase(repo domain.Repository) CreateUseCase {
	return &createUseCase{
		repo: repo,
	}
}

func (u *createUseCase) Execute(ctx context.Context, model *domain.Category) (*domain.Category, []response.ErrorField, error) {
	existingModel, err := u.repo.GetByUniqueFields(ctx, model.Name, model.Slug)
	if err != nil {
		return nil, nil, err
	}
	if existingModel != nil {
		var validationErrors []response.ErrorField
		if existingModel.Name == model.Name {
			validationErrors = append(validationErrors, response.NewErrorField("name", string(response.NotUnique)))
		}
		if existingModel.Slug == model.Slug {
			validationErrors = append(validationErrors, response.NewErrorField("slug", string(response.NotUnique)))
		}
		return nil, validationErrors, nil
	}

	createdModel, err := u.repo.Create(ctx, model)
	if err != nil {
		return nil, nil, err
	}

	return createdModel, nil, nil
}
