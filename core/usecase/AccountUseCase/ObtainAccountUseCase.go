package AccountUseCase

import (
	"context"

	"github.com/sabidos/core/entity"
	"go.mongodb.org/mongo-driver/bson"
)

type ObtainAccountUseCase struct {
	accountRepository entity.AccountDataProvider
}

func NewObtainAccountUsecase(a entity.AccountDataProvider) entity.ObtainAccountUseCase {
	return &ObtainAccountUseCase{
		accountRepository: a,
	}
}

func (a *ObtainAccountUseCase) Get(c context.Context, filter bson.M) (acc entity.Account, err error) {
	acc, err = a.accountRepository.Get(c, filter)
	if err != nil {
		return acc, err
	}
	return acc, err
}
