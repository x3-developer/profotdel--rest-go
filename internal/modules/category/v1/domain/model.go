package domain

import "time"

type Category struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Name      string    `gorm:"type:varchar(255);not null;unique" json:"name"`
	Slug      string    `gorm:"type:varchar(255);not null;unique" json:"slug"`
}
