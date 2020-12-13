package AccountUseCase

import (
	"context"
	"errors"

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

	if acc.Avatar.Id == 0 || acc.IsAnonymous {
		avatarCount, err := a.avatarDataProvider.Count(c, bson.M{})
		if err != nil {
			return account, err
		}
		acc.SetRandomAvatar(avatarCount)
	}

	avatarFilter := bson.M{"id": acc.Avatar.Id}

	avatar, _ := a.avatarDataProvider.FindOne(c, avatarFilter)

	if avatar.Id == 0 {
		return account, errors.New("Avatar does not exists")
	}

	acc.SetAvatar(avatar)
	acc.SetXpFactor(3) // Every new account will be receiving fixed xpFactor

	// Check if is an Anonymous user and complete acc info 
	acc.CompleteAccountIfAnonymous()
	
	err = a.accountRepository.Insert(c, acc)

	if err != nil {
		return account, err
	}

	return acc, err
}
