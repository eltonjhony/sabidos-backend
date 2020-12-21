package entity

import (
	"context"
)

type Avatar struct {
	Id       int    `json:"id"`
	ImageUrl string `json:"imageUrl"`
}

type AvatarDataProvider interface {
	GetAll(ctx context.Context) (avatar []Avatar, err error)
	Count(ctx context.Context) (itemCount int64, err error)
	Insert(ctx context.Context, avatar Avatar) error
	FindById(ctx context.Context, id int) (avatar Avatar, err error)
}

type ObtainAvatarUseCase interface {
	GetAll(ctx context.Context) (avatar []Avatar, err error)
}
