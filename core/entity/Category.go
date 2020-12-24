package entity

import (
	"context"
)

type Category struct {
	Id          int    `json:"id"`
	ImageUrl    string `json:"imageUrl"`
	Description string `json:"description"`
	IconUrl     string `json:"iconUrl"`
}

type CategoryDataProvider interface {
	GetAll(ctx context.Context) (category []Category, err error)
	Insert(ctx context.Context, category Category) error
	FindById(ctx context.Context, id int) (category Category, err error)
}
