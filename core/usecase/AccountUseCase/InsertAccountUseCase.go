package AccountUseCase

import (
	"context"
	"errors"

	"github.com/sabidos/core/entity"
	"go.mongodb.org/mongo-driver/bson"
)

type InserAccountUseCase struct {
	accountRepository entity.AccountDataProvider
}

func NewInsertAccountUsecase(a entity.AccountDataProvider) entity.InsertAccountUseCase {
	return &InserAccountUseCase{
		accountRepository: a,
	}
}

func (a *InserAccountUseCase) Insert(c context.Context, acc entity.Account) (err error) {

	bfilter := bson.M{"$or": []bson.M{bson.M{"nickname": acc.NickName}, bson.M{"uid": acc.Uid}}}

	if account, _ := a.accountRepository.Get(c, bfilter); len(account.NickName) > 0 {
		return errors.New("Account already exists")

	}

	err = a.accountRepository.Insert(c, acc)

	if err != nil {
		return err
	}

	return
}
