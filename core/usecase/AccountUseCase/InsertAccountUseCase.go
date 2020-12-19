package AccountUseCase

import (
	"context"
	"errors"
	"math/rand"
	"time"
	"fmt"

	"github.com/sabidos/core/entity"
	"go.mongodb.org/mongo-driver/bson"
)

type InserAccountUseCase struct {
	accountRepository  entity.AccountDataProvider
	avatarDataProvider entity.AvatarDataProvider
}

func NewInsertAccountUsecase(acc entity.AccountDataProvider, avatar entity.AvatarDataProvider) entity.InsertAccountUseCase {
	return &InserAccountUseCase{
		accountRepository:  acc,
		avatarDataProvider: avatar,
	}
}

func (a *InserAccountUseCase) Insert(c context.Context, acc entity.Account) (account entity.Account, err error) {

	bfilter := bson.M{"$or": []bson.M{bson.M{"nickname": acc.NickName}, bson.M{"uid": acc.Uid}}}

	if account, _ := a.accountRepository.Get(c, bfilter); len(account.NickName) > 0 {
		return account, errors.New("Account already exists")
	}

	account = acc
	
	avatar, err := a.getAccountAvatar(c, acc)
	if err != nil {
		return account, err
	}

	account.SetAvatar(avatar)
	account.SetTotalAnswered(0)
	account.SetTotalHits(0)
	account.SetXpFactor(3)
	account.SetReputation(1, 0)
	account.CompleteAccountIfAnonymous()
	
	err = a.accountRepository.Insert(c, account)

	if err != nil {
		return account, err
	}

	return account, err
}

func (a *InserAccountUseCase) getAccountAvatar(c context.Context, acc entity.Account) (res entity.Avatar, err error)  {
	var result entity.Avatar
	if acc.Avatar.Id == 0 || acc.IsAnonymous {
		avatar, err := a.getRandomAvatar(c)
		if err != nil {
			return res, err
		}

		result = avatar
	} else {
		avatarFilter := bson.M{"id": acc.Avatar.Id}
		avatar, err := a.avatarDataProvider.FindOne(c, avatarFilter)
		if avatar.Id == 0 || err != nil {
			return res, err
		}

		result = avatar
	}

	return result, err
}

func (a *InserAccountUseCase) getRandomAvatar(c context.Context) (res entity.Avatar, err error) {
	avatarCount, err := a.avatarDataProvider.Count(c, bson.M{})
	if err != nil {
		return res, err
	}
	
	rand.Seed(time.Now().UnixNano())
	min := 1
	max := int(avatarCount)
	
	// Get Random avatarId in Range
	randomAvatarId := rand.Intn(max - min + 1) + min
	fmt.Println("Random avatar id ", randomAvatarId)

	avatarFilter := bson.M{"id": randomAvatarId}
	return a.avatarDataProvider.FindOne(c, avatarFilter)

}