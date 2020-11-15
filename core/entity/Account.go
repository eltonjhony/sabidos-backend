package entity

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

type Account struct {
	Uid           string     `json:"uid"`
	Name          string     `json:"name"`
	NickName      string     `json:"nickname"`
	Avatar        Avatar     `json:"avatar"`
	Reputation    Reputation `json:"reputation"`
	TotalAnswered int        `json:"totalAnswered"`
	TotalHits     int        `json:"totalHits"`
	Email         string     `json:"email"`
	IsAnonymous   bool       `json:"isAnonymous"`
	Phone         string     `json:"phone"`
}

func (acc *Account) SetAvatar(avatar Avatar) {
	acc.Avatar = avatar
}

type ObtainAccountUseCase interface {
	Get(ctx context.Context, filter bson.M) (acc Account, err error)
}

type InsertAccountUseCase interface {
	Insert(ctx context.Context, acc Account) (account Account, err error)
}

type AccountDataProvider interface {
	Get(ctx context.Context, filter bson.M) (account Account, err error)
	Insert(ctx context.Context, acc Account) error
}
