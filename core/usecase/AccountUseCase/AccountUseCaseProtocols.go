package AccountUseCase

import (
	"context"

	"github.com/sabidos/core/entity"
	"github.com/sabidos/entrypoint/model"
)

type ObtainAccountUseCaseProtocol interface {
	GetByNickname(ctx context.Context, nickname string) (acc entity.Account, err error)
	GetByUid(ctx context.Context, uid string) (acc entity.Account, err error)
}

type InsertAccountUseCaseProtocol interface {
	Insert(ctx context.Context, model model.AccountModel) (account entity.Account, err error)
}

type UpdateAccountUseCaseProtocol interface {
	Update(ctx context.Context, uid string, model model.UpdateAccountModel) (account entity.Account, err error)
}

type ValidateAccountUseCaseProtocol interface {
	Validate(c context.Context, nickname string, uid string) error
}
