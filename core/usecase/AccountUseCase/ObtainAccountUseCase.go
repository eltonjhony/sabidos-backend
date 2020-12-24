package AccountUseCase

import (
	"context"
	"fmt"

	"github.com/sabidos/core/entity"
)

type ObtainAccountUseCase struct {
	accountRepository entity.AccountDataProvider
}

func NewObtainAccountUsecase(a entity.AccountDataProvider) ObtainAccountUseCaseProtocol {
	return &ObtainAccountUseCase{
		accountRepository: a,
	}
}

func (a *ObtainAccountUseCase) GetByNickname(c context.Context, nickname string) (acc entity.Account, err error) {

	acc, err = a.accountRepository.GetByNickname(c, nickname)
	if err != nil {
		fmt.Printf("Error %s ", err)
		return acc, err
	}
	return acc, err
}

func (a *ObtainAccountUseCase) GetByUid(c context.Context, uid string) (acc entity.Account, err error) {

	acc, err = a.accountRepository.GetByUid(c, uid)
	if err != nil {
		fmt.Printf("Error %s ", err)
		return acc, err
	}
	return acc, err
}
