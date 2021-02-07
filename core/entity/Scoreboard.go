package entity

import "context"

type ScoreboardDataProvider interface {
	GetByIdentifier(ctx context.Context, nickname string, scoreEndTimestamp int) (scoreboard Scoreboard, err error)
	Insert(ctx context.Context, scoreboard Scoreboard) error
	Update(ctx context.Context, scoreboard Scoreboard) error
}

type Scoreboard struct {
	Nickname          string `json:"nickname"`
	HitsAmount        int    `json:"hitsAmount"`
	ScoreEndTimestamp int    `json:"scoreEndTimestamp"`
}

func (scoreboard *Scoreboard) IncreaseHitsAmount() {
	scoreboard.HitsAmount += 1
}
