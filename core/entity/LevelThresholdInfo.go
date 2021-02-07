package entity

import "context"

type LevelThresholdInfoDataProvider interface {
	GetAll(ctx context.Context) (levelThresholdInfo []LevelThresholdInfo, err error)
	Insert(ctx context.Context, levelThresholdInfo LevelThresholdInfo) error
}

type LevelThresholdInfo struct {
	Level                 int `json:"level"`
	HitsPerStar           int `json:"hitsPerStar"`
	TotalStarsToNextLevel int `json:totalStarsToNextLevel`
}
