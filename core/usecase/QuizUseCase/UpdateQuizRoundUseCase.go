package QuizUseCase

import (
	"context"
	"fmt"

	"github.com/sabidos/core/entity"
)

type UpdateQuizRoundUseCase struct {
	accountRepository entity.AccountDataProvider
}

func NewUpdateQuizRoundUseCase(a entity.AccountDataProvider) entity.UpdateQuizRoundUseCase {
	return &UpdateQuizRoundUseCase{
		accountRepository: a,
	}
}

func (a *UpdateQuizRoundUseCase) UpdateQuizRoundValues(ctx context.Context, nickname string, accumulateXp int) (err error) {
	fmt.Printf("Starting updating Quiz Round values for nickname %s", nickname)

	account, err := a.accountRepository.GetByNickname(ctx, nickname)
	if err != nil {
		return err
	}

	account.AddAccumulateXp(accumulateXp)
	err = a.accountRepository.Update(ctx, account)
	if err != nil {
		return err
	}

	return err
}
