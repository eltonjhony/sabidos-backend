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
	TotalAnswered string     `json:"totalAnswered"`
	TotalHits     string     `json:"totalHits"`
	Email         string     `json:"email"`
	IsAnonymous   bool       `json:"isAnonymous"`
	Phone         string     `json:"phone"`
}

type ObtainAccountUseCase interface {
	Get(ctx context.Context, filter bson.M) (acc Account, err error)
}

type InsertAccountUseCase interface {
	Insert(ctx context.Context, acc Account) error
}

type AccountDataProvider interface {
	Get(ctx context.Context, filter bson.M) (acc Account, err error)
	Insert(ctx context.Context, acc Account) error
}
