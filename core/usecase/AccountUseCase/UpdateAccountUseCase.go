package AccountUseCase

import (
	"context"
	"errors"

	"github.com/sabidos/core/entity"
	"go.mongodb.org/mongo-driver/bson"
)

type UpdateAccountUseCase struct {
	accountRepository entity.AccountDataProvider
}

func NewUpdateAccountUsecase(acc entity.AccountDataProvider) entity.UpdateAccountUseCase {
	return &UpdateAccountUseCase{
		accountRepository: acc,
	}
}

func (a *UpdateAccountUseCase) Update(c context.Context, uid string, acc entity.Account) (entity.Account, error) {

	bfilter := bson.M{"uid": uid}
	account, err := a.accountRepository.Get(c, bfilter)

	if err != nil || len(account.Uid) <= 0 {
		return account, errors.New("Account not found")
	}

	if len(acc.Name) > 0 {
		account.Name = acc.Name
	}

	if len(acc.Email) > 0 {
		account.Email = acc.Email
	}

	if len(acc.Phone) > 0 {
		account.Phone = acc.Phone
	}

	account.IsAnonymous = acc.IsAnonymous

	err = a.accountRepository.Update(c, account)

	if err != nil {
		return account, err
	}

	return account, err

}
