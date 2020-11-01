package entity

import "context"

type Account struct {
	Id            int        `json:"id"`
	Name          string     `json:"name"`
	NickName      string     `json:"nickname"`
	Avatar        Avatar     `json:"avatar"`
	Reputation    Reputation `json:"reputation"`
	TotalAnswered string     `json:"totalAnswered"`
	TotalHits     string     `json:"totalHits"`
}

func (acc *Account) SetId(id int) {
	acc.Id = id
}

type ObtainAccountUseCase interface {
	Get(ctx context.Context, id string) (acc Account, err error)
}

type InsertAccountUseCase interface {
	Insert(ctx context.Context, acc Account) error
}

type AccountDataProvider interface {
	Get(ctx context.Context, id string) (acc Account, err error)
	Insert(ctx context.Context, acc Account) error
}
