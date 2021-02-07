package entity

import (
	"context"
	"fmt"
	"time"
)

type AccountDataProvider interface {
	GetByIdentifier(ctx context.Context, nickname string, uid string) (account Account, err error)
	GetByNickname(ctx context.Context, nickname string) (account Account, err error)
	GetByUid(ctx context.Context, uid string) (account Account, err error)
	Insert(ctx context.Context, acc Account) error
	Update(ctx context.Context, acc Account) error
}

type Account struct {
	Uid                    string     `json:"uid"`
	Name                   string     `json:"name"`
	NickName               string     `json:"nickname"`
	Avatar                 Avatar     `json:"avatar"`
	Reputation             Reputation `json:"reputation"`
	XpFactor               int        `json:"xpFactor"`
	TotalAnswered          int        `json:"totalAnswered"`
	TotalHits              int        `json:"totalHits"`
	Email                  string     `json:"email"`
	IsAnonymous            bool       `json:"isAnonymous"`
	Phone                  string     `json:"phone"`
	AccumulateXp           int        `json:"accumulateXp"`
	AccumulateResponseTime int        `json:"accumulateResponseTime"`
	AnsweredQuiz           []string   `json:"answeredQuiz"`
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

func (acc *Account) IncreaseTotalAnswered() {
	acc.TotalAnswered += 1
}

func (acc *Account) IncreaseTotalHits() {
	acc.TotalHits += 1
	acc.AccumulateXp += acc.XpFactor
}

func (acc *Account) AddAccumulateResponseTime(responseTime int) {
	acc.AccumulateResponseTime += responseTime
}

func (acc *Account) AddAnsweredQuiz(quizId string) {
	acc.AnsweredQuiz = append(acc.AnsweredQuiz, quizId)
}

func (acc *Account) CompleteAccountIfAnonymous() {
	if acc.IsAnonymous {
		timestamp := time.Now().UnixNano() / int64(time.Millisecond)
		acc.Name = fmt.Sprintf("%s #%d", "Player", timestamp)
		acc.NickName = fmt.Sprintf("%d", timestamp)
	}
}
