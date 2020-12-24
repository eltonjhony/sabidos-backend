package AccountUseCase

import (
	"context"
	"errors"

	"github.com/sabidos/core/entity"
	"github.com/sabidos/entrypoint/model"
)

type UpdateAccountUseCase struct {
	accountRepository entity.AccountDataProvider
}

func NewUpdateAccountUsecase(acc entity.AccountDataProvider) UpdateAccountUseCaseProtocol {
	return &UpdateAccountUseCase{
		accountRepository: acc,
	}
}

func (a *UpdateAccountUseCase) Update(c context.Context, uid string, model model.UpdateAccountModel) (entity.Account, error) {

	account, err := a.accountRepository.GetByUid(c, uid)

	if err != nil || len(account.Uid) <= 0 {
		return account, errors.New("Account not found")
	}

	if len(model.Name) > 0 {
		account.Name = model.Name
	}

	if len(model.Email) > 0 {
		account.Email = model.Email
	}

	if len(model.Phone) > 0 {
		account.Phone = model.Phone
	}

	account.IsAnonymous = model.IsAnonymous

	err = a.accountRepository.Update(c, account)

	if err != nil {
		return account, err
	}

	return account, err

}
