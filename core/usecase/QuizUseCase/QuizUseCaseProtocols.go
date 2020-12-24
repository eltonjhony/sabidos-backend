package QuizUseCase

import (
	"context"

	"github.com/sabidos/core/entity"
	"github.com/sabidos/entrypoint/model"
)

type ObtainQuizUseCaseProtocol interface {
	ObtainQuizRoundFor(ctx context.Context, nickname string, categoryId string) ([]entity.Quiz, error)
}

type UpdateQuizRoundUseCaseProtocol interface {
	UpdateQuizRoundValues(ctx context.Context, requestModel model.PostRoundModel) (err error)
}
