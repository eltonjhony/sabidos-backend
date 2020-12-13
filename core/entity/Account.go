package entity

import (
	"context"
	"time"
	"math/rand"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
)

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
}

func (acc *Account) SetAvatar(avatar Avatar) {
	acc.Avatar = avatar
}

func (acc *Account) SetXpFactor(xpFactor int) {
	acc.XpFactor = xpFactor
}

func (acc *Account) SetRandomAvatar() {
	rand.Seed(time.Now().UnixNano())
	min := 1
	max := 15
	// Get Random avatarId in Range
	randomAvatarId := rand.Intn(max - min + 1) + min
	fmt.Println("%s #%d", "Random avatar id", randomAvatarId)
	acc.SetAvatar(Avatar{randomAvatarId, ""})
}

func (acc *Account) CompleteAccountIfAnonymous() {
	if acc.IsAnonymous {
		timestamp := time.Now().UnixNano() / int64(time.Millisecond)
		acc.Name = fmt.Sprintf("%s #%d", "Player", timestamp)
		acc.NickName = fmt.Sprintf("%d", timestamp)
	}
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
