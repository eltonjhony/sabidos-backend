package AccountUseCase

import (
	"context"
	"fmt"

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

func (a *ObtainAccountUseCase) GetByNickname(c context.Context, nickname string) (acc entity.Account, err error) {

	bfilter := bson.M{"nickname": nickname}

	acc, err = a.accountRepository.Get(c, bfilter)
	if err != nil {
		fmt.Printf("Error %s ", err)
		return acc, err
	}
	return acc, err
}

func (a *ObtainAccountUseCase) GetByUid(c context.Context, uid string) (acc entity.Account, err error) {

	bfilter := bson.M{"uid": uid}

	acc, err = a.accountRepository.Get(c, bfilter)
	if err != nil {
		fmt.Printf("Error %s ", err)
		return acc, err
	}
	return acc, err
}
