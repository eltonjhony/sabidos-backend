package AccountUseCase

import (
	"context"

	"github.com/sabidos/core/entity"
)

type ValidateAccountUseCase struct {
	accountRepository entity.AccountDataProvider
}

func NewValidateAccountUsecase(acc entity.AccountDataProvider) entity.ValidateAccountUseCase {
	return &ValidateAccountUseCase{
		accountRepository: acc,
	}
}

func (a *ValidateAccountUseCase) Validate(c context.Context, nickname string, uid string) error {
	_, err := a.accountRepository.GetByIdentifier(c, nickname, uid)
	return err
}
