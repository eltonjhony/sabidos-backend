package QuizUseCase

import (
	"context"
	"fmt"

	"github.com/sabidos/core/entity"
	"github.com/sabidos/entrypoint/model"
)

type UpdateQuizRoundUseCase struct {
	accountRepository entity.AccountDataProvider
}

func NewUpdateQuizRoundUseCase(a entity.AccountDataProvider) UpdateQuizRoundUseCaseProtocol {
	return &UpdateQuizRoundUseCase{
		accountRepository: a,
	}
}

func (a *UpdateQuizRoundUseCase) UpdateQuizRoundValues(ctx context.Context, requestModel model.PostRoundModel) (err error) {
	fmt.Printf("Starting updating Quiz Round values for nickname %s", requestModel.NickName)

	account, err := a.accountRepository.GetByNickname(ctx, requestModel.NickName)
	if err != nil {
		return err
	}

	account.AddAccumulateXp(requestModel.AccumulateXp)
	err = a.accountRepository.Update(ctx, account)
	if err != nil {
		return err
	}

	return err
}
