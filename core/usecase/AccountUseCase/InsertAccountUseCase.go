package AccountUseCase

import (
	"context"

	"github.com/sabidos/core/entity"
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
	err = a.accountRepository.Insert(c, acc)
	if err != nil {
		return err
	}
	return
}
