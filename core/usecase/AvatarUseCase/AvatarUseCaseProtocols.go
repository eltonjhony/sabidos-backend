package AvatarUseCase

import (
	"context"

	"github.com/sabidos/core/entity"
)

type ObtainAvatarUseCaseProtocol interface {
	GetAll(ctx context.Context) (avatar []entity.Avatar, err error)
}
