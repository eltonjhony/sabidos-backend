package QuizUseCase

import (
	"context"
	"fmt"
	"errors"
	"strconv"

	"github.com/sabidos/core/entity"
	"go.mongodb.org/mongo-driver/bson"
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
	
	// Filter params
	bfilter := bson.M{"category.id": catId}
	recordsLimit := int64(10)

	quizRound, err := a.quizRepository.Get(ctx, bfilter, recordsLimit)
	return quizRound, err
}
