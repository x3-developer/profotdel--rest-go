package dto

type CreateDTO struct {
	Name string `json:"name" validate:"required,min=3,max=64"`
	Slug string `json:"slug" validate:"required,min=3,max=64"`
}
