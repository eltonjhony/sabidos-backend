package RankingUseCase

import (
	"context"

	"github.com/sabidos/core/entity"
)

type RankingUseCase struct {
	rankingRepo entity.RankingRepository
}

func NewRankingUsecase(a entity.RankingRepository) entity.RankingUseCase {
	return &RankingUseCase{
		rankingRepo: a,
	}
}

func (a *RankingUseCase) Fetch(c context.Context) (res []entity.Ranking, err error) {
	res, err = a.rankingRepo.Fetch(c)
	if err != nil {
		return nil, err
	}
	return
}
