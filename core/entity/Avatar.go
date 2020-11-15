package entity

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

type Avatar struct {
	Id       int    `json:"id"`
	ImageUrl string `json:"imageUrl"`
}

type AvatarDataProvider interface {
	Get(ctx context.Context, filter bson.M) (avatar []Avatar, err error)
	Insert(ctx context.Context, avatar Avatar) error
	FindOne(ctx context.Context, filter bson.M) (avatar Avatar, err error)
}

type ObtainAvatarUseCase interface {
	Get(ctx context.Context, filter bson.M) (avatar []Avatar, err error)
}
