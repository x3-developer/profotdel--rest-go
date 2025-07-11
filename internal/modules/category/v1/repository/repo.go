package repository

import (
	"context"
	"profotdel-rest/internal/modules/category/v1/domain"
	"profotdel-rest/pkg/database"
)

type Repository interface {
	Create(ctx context.Context, model *domain.Category) (*domain.Category, error)
	GetAll(ctx context.Context) ([]*domain.Category, error)
	GetById(ctx context.Context, id uint) (*domain.Category, error)
	Update(ctx context.Context, model *domain.Category) (*domain.Category, error)
	Delete(ctx context.Context, id uint) error
	GetByUniqueFields(ctx context.Context, name, slug string) (*domain.Category, error)
}

type repository struct {
	DB *database.DB
}

func NewRepository(db *database.DB) Repository {
	return &repository{
		DB: db,
	}
}

func (r *repository) Create(ctx context.Context, model *domain.Category) (*domain.Category, error) {
	result := r.DB.WithContext(ctx).Create(&model)
	if result.Error != nil {
		return nil, result.Error
	}

	return model, nil
}

func (r *repository) GetAll(ctx context.Context) ([]*domain.Category, error) {
	var models []*domain.Category

	result := r.DB.WithContext(ctx).Find(&models)
	if result.Error != nil {
		return nil, result.Error
	}

	return models, nil
}

func (r *repository) GetById(ctx context.Context, id uint) (*domain.Category, error) {
	var model *domain.Category

	result := r.DB.WithContext(ctx).First(&model, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return model, nil
}

func (r *repository) Update(ctx context.Context, model *domain.Category) (*domain.Category, error) {
	result := r.DB.WithContext(ctx).Save(&model)
	if result.Error != nil {
		return nil, result.Error
	}

	return model, nil
}

func (r *repository) Delete(ctx context.Context, id uint) error {
	var model *domain.Category

	result := r.DB.WithContext(ctx).Delete(model, id)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *repository) GetByUniqueFields(ctx context.Context, name, slug string) (*domain.Category, error) {
	var model *domain.Category

	result := r.DB.WithContext(ctx).Where("name = ? OR slug = ?", name, slug).First(&model)
	if result.Error != nil {
		return nil, result.Error
	}

	return model, nil
}
