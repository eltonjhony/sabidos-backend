package AccountUseCase

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/sabidos/core/entity"
	"github.com/sabidos/entrypoint/model"
)

type InsertAccountUseCase struct {
	accountRepository  entity.AccountDataProvider
	avatarDataProvider entity.AvatarDataProvider
}

func NewInsertAccountUsecase(acc entity.AccountDataProvider, avatar entity.AvatarDataProvider) InsertAccountUseCaseProtocol {
	return &InsertAccountUseCase{
		accountRepository:  acc,
		avatarDataProvider: avatar,
	}
}

func (a *InsertAccountUseCase) Insert(c context.Context, model model.AccountModel) (account entity.Account, err error) {

	if account, _ := a.accountRepository.GetByIdentifier(c, model.NickName, model.Uid); len(account.NickName) > 0 {
		return account, errors.New("Account already exists")
	}

	account = entity.Account{
		Uid:      model.Uid,
		Name:     model.Name,
		NickName: model.NickName,
		Avatar: entity.Avatar{
			Id: model.DefaultAvatarId,
		},
		Email:       model.Email,
		IsAnonymous: model.IsAnonymous,
		Phone:       model.Phone,
	}

	avatar, err := a.getAccountAvatar(c, account)
	if err != nil {
		return account, err
	}

	account.SetAvatar(avatar)
	account.SetTotalAnswered(0)
	account.SetTotalHits(0)
	account.SetXpFactor(3)
	account.SetReputation(1, 0)
	account.AddAccumulateXp(0)
	account.CompleteAccountIfAnonymous()

	err = a.accountRepository.Insert(c, account)

	if err != nil {
		return account, err
	}

	return account, err
}

func (a *InsertAccountUseCase) getAccountAvatar(c context.Context, acc entity.Account) (res entity.Avatar, err error) {
	var result entity.Avatar
	if acc.Avatar.Id == 0 || acc.IsAnonymous {
		avatar, err := a.getRandomAvatar(c)
		if err != nil {
			return res, err
		}

		result = avatar
	} else {
		avatar, err := a.avatarDataProvider.FindById(c, acc.Avatar.Id)
		if avatar.Id == 0 || err != nil {
			return res, err
		}

		result = avatar
	}

	return result, err
}

func (a *InsertAccountUseCase) getRandomAvatar(c context.Context) (res entity.Avatar, err error) {
	avatarCount, err := a.avatarDataProvider.Count(c)
	if err != nil {
		return res, err
	}

	rand.Seed(time.Now().UnixNano())
	min := 1
	max := int(avatarCount)

	// Get Random avatarId in Range
	randomAvatarId := rand.Intn(max-min+1) + min
	fmt.Println("Random avatar id ", randomAvatarId)

	return a.avatarDataProvider.FindById(c, randomAvatarId)

}
