package QuizUseCase

import (
	"context"

	"github.com/sabidos/core/entity"
	"github.com/sabidos/entrypoint/model"
)

type ObtainQuizUseCaseProtocol interface {
	ObtainQuizRoundFor(ctx context.Context, nickname string, categoryId string) ([]entity.Quiz, error)
}

type UpdateQuizAccountValuesUseCaseProtocol interface {
	UpdateQuizAccountValues(ctx context.Context, requestModel model.PostQuizModel) (levelHasBeenUp bool, starHasBeenUp bool, err error)
}
