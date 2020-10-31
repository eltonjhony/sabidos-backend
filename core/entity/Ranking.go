package entity

import "context"

type Ranking struct {
	Name string `json:"name"`
}

type Filter struct {
	Name string `json:"name"`
}

type RankingUseCase interface {
	Fetch(ctx context.Context) ([]Ranking, error)
}

type RankingRepository interface {
	Fetch(ctx context.Context) (res []Ranking, err error)
}
