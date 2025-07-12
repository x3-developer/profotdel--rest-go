package mapper

import (
	"profotdel-rest/internal/modules/category/v1/application/dto"
	"profotdel-rest/internal/modules/category/v1/domain"
)

func ToResponseDTOFromModel(model *domain.Category) *dto.ResponseDTO {
	if model == nil {
		return nil
	}

	return &dto.ResponseDTO{
		Id:        model.ID,
		Name:      model.Name,
		Slug:      model.Slug,
		CreatedAt: model.CreatedAt,
		UpdatedAt: model.UpdatedAt,
	}
}

func ToModelFromCreateDTO(createDto *dto.CreateDTO) *domain.Category {
	if createDto == nil {
		return nil
	}

	return &domain.Category{
		Name: createDto.Name,
		Slug: createDto.Slug,
	}
}
