package entity

import "context"

type Account struct {
	Name          string     `json:"name"`
	NickName      string     `json:"nickname"`
	Avatar        Avatar     `json:"avatar"`
	Reputation    Reputation `json:"reputation"`
	TotalAnswered string     `json:"totalAnswered"`
	TotalHits     string     `json:"totalHits"`
}

type ObtainAccountUseCase interface {
	Fetch(ctx context.Context) ([]Account, error)
}

type AccountDataProvider interface {
	Fetch(ctx context.Context) (res []Account, err error)
	Insert(acc Account, ctx context.Context)
}
