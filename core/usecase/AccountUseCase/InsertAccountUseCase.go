package AccountUseCase

import (
	"context"
	"errors"
	"fmt"

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

	bfilter := bson.M{"nickname": acc.NickName}

	acc, err = a.accountRepository.Get(c, bfilter)
	fmt.Printf("%+v\n", acc)

	if len(acc.NickName) > 0 {
		return errors.New("Account already exists")
	}

	err = a.accountRepository.Insert(c, acc)

	if err != nil {
		return err
	}

	return
}
