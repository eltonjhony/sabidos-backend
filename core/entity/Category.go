package entity

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

type Category struct {
	Id          int    `json:"id"`
	ImageUrl    string `json:"imageUrl"`
	Description string `json:"description"`
	IconUrl     string `json:"iconUrl"`
}

type CategoryDataProvider interface {
	Get(ctx context.Context, filter bson.M) (category []Category, err error)
	Insert(ctx context.Context, category Category) error
	FindOne(ctx context.Context, filter bson.M) (category Category, err error)
}

type ObtainCategoryUseCase interface {
	Get(ctx context.Context, filter bson.M) (category []Category, err error)
}
