package entity

import (
	"context"
	"fmt"
	"time"
)

type ObtainAccountUseCase interface {
	GetByNickname(ctx context.Context, nickname string) (acc Account, err error)
	GetByUid(ctx context.Context, uid string) (acc Account, err error)
}

type InsertAccountUseCase interface {
	Insert(ctx context.Context, acc Account) (account Account, err error)
}

type UpdateAccountUseCase interface {
	Update(ctx context.Context, uid string, acc Account) (account Account, err error)
}

type ValidateAccountUseCase interface {
	Validate(c context.Context, nickname string, uid string) error
}

type AccountDataProvider interface {
	GetByIdentifier(ctx context.Context, nickname string, uid string) (account Account, err error)
	GetByNickname(ctx context.Context, nickname string) (account Account, err error)
	GetByUid(ctx context.Context, uid string) (account Account, err error)
	Insert(ctx context.Context, acc Account) error
	Update(ctx context.Context, acc Account) error
}

type Account struct {
	Uid           string     `json:"uid"`
	Name          string     `json:"name"`
	NickName      string     `json:"nickname"`
	Avatar        Avatar     `json:"avatar"`
	Reputation    Reputation `json:"reputation"`
	XpFactor      int        `json:"xpFactor"`
	TotalAnswered int        `json:"totalAnswered"`
	TotalHits     int        `json:"totalHits"`
	Email         string     `json:"email"`
	IsAnonymous   bool       `json:"isAnonymous"`
	Phone         string     `json:"phone"`
	AccumulateXp  int        `json:"accumulateXp"`
}

func (acc *Account) SetAvatar(avatar Avatar) {
	acc.Avatar = avatar
}

func (acc *Account) SetXpFactor(xpFactor int) {
	acc.XpFactor = xpFactor
}

func (acc *Account) SetReputation(level int, stars int) {
	acc.Reputation = Reputation{
		Level: level,
		Stars: stars,
	}
}

func (acc *Account) SetTotalAnswered(totalAnswered int) {
	acc.TotalAnswered = totalAnswered
}

func (acc *Account) SetTotalHits(totalHits int) {
	acc.TotalHits = totalHits
}

func (acc *Account) AddAccumulateXp(accumulateXp int) {
	acc.AccumulateXp += accumulateXp
}

func (acc *Account) CompleteAccountIfAnonymous() {
	if acc.IsAnonymous {
		timestamp := time.Now().UnixNano() / int64(time.Millisecond)
		acc.Name = fmt.Sprintf("%s #%d", "Player", timestamp)
		acc.NickName = fmt.Sprintf("%d", timestamp)
	}
}
