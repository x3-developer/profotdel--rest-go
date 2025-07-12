package domain

import "context"

type Repository interface {
	Create(ctx context.Context, model *Category) (*Category, error)
	GetAll(ctx context.Context) ([]*Category, error)
	GetById(ctx context.Context, id uint) (*Category, error)
	Update(ctx context.Context, model *Category) (*Category, error)
	Delete(ctx context.Context, id uint) error
	GetByUniqueFields(ctx context.Context, name, slug string) (*Category, error)
}
