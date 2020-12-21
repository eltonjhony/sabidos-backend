package AccountUseCase

import (
	"context"

	"github.com/sabidos/core/entity"
	"go.mongodb.org/mongo-driver/bson"
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
	bfilter := bson.M{"$or": []bson.M{bson.M{"nickname": nickname}, bson.M{"uid": uid}}}
	_, err := a.accountRepository.Get(c, bfilter)
	return err
}
