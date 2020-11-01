package AccountUseCase

import (
	"context"

	"github.com/sabidos/core/entity"
)

type ObtainAccountUseCase struct {
	accountRepository entity.AccountDataProvider
}

func NewObtainAccountUsecase(a entity.AccountDataProvider) entity.ObtainAccountUseCase {
	return &ObtainAccountUseCase{
		accountRepository: a,
	}
}

func (a *ObtainAccountUseCase) Get(c context.Context, id string) (acc entity.Account, err error) {
	acc, err = a.accountRepository.Get(c, id)
	if err != nil {
		return acc, err
	}
	return acc, err
}
