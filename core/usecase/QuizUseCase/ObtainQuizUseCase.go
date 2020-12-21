package QuizUseCase

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"github.com/sabidos/core/entity"
)

type ObtainQuizUseCase struct {
	quizRepository entity.QuizDataProvider
}

func NewObtainQuizUsecase(a entity.QuizDataProvider) entity.ObtainQuizUseCase {
	return &ObtainQuizUseCase{
		quizRepository: a,
	}
}

func (a *ObtainQuizUseCase) ObtainQuizRoundFor(ctx context.Context, nickname string, categoryId string) (res []entity.Quiz, err error) {
	fmt.Printf("Starting obtaining Quiz Round for nickname %s and category %s", nickname, categoryId)

	if len(categoryId) == 0 {
		return nil, errors.New("CategoryId is a required param")
	}

	catId, err := strconv.Atoi(categoryId)
	if err != nil {
		return nil, errors.New("Error converting category id")
	}

	quizRound, err := a.quizRepository.GetByCategory(ctx, catId, 10)
	return quizRound, err
}
