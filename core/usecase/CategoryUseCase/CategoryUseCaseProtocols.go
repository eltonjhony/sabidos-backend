package CategoryUseCase

import (
	"context"

	"github.com/sabidos/core/entity"
)

type ObtainCategoryUseCaseProtocol interface {
	GetAll(ctx context.Context) (category []entity.Category, err error)
}
