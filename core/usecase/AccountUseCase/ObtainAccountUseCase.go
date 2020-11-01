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

func (a *ObtainAccountUseCase) Fetch(c context.Context) (res []entity.Account, err error) {
	res, err = a.accountRepository.Fetch(c)
	if err != nil {
		return nil, err
	}
	return
}
