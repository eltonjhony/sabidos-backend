package RankingUseCase

import (
	"context"

	"github.com/sabidos/core/entity"
)

type RankingUseCaseProtocol interface {
	Fetch(ctx context.Context) ([]entity.Ranking, error)
}
