package entity

import "context"

type Ranking struct {
	Name string `json:"name"`
}

type Filter struct {
	Name string `json:"name"`
}

type RankingDataProvider interface {
	Fetch(ctx context.Context) (res []Ranking, err error)
}
