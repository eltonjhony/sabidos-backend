package QuizUseCase

import (
	"context"
	"fmt"

	"github.com/sabidos/core/entity"
	"github.com/sabidos/entrypoint/model"
	"github.com/sabidos/utils"
)

type UpdateQuizAccountValuesUseCase struct {
	accountRepository        entity.AccountDataProvider
	scoreboardRepository     entity.ScoreboardDataProvider
	levelThresholdRepository entity.LevelThresholdInfoDataProvider
}

func NewUpdateQuizAccountValuesUseCase(a entity.AccountDataProvider, b entity.ScoreboardDataProvider, c entity.LevelThresholdInfoDataProvider) UpdateQuizAccountValuesUseCaseProtocol {
	return &UpdateQuizAccountValuesUseCase{
		accountRepository:        a,
		scoreboardRepository:     b,
		levelThresholdRepository: c,
	}
}

func (self *UpdateQuizAccountValuesUseCase) UpdateQuizAccountValues(ctx context.Context, requestModel model.PostQuizModel) (levelHasBeenUp bool, starHasBeenUp bool, err error) {
	fmt.Printf("Starting updating account values per Quiz for nickname %s", requestModel.NickName)

	account, err := self.accountRepository.GetByNickname(ctx, requestModel.NickName)
	if err != nil {
		fmt.Printf("Error fetching account by nickname %s", requestModel.NickName)
		return levelHasBeenUp, starHasBeenUp, err
	}

	account.IncreaseTotalAnswered()

	if requestModel.Alternative.IsCorrect {
		account.IncreaseTotalHits()
		scoreboard, err := self.updateScoreboard(ctx, account)
		if err != nil {
			fmt.Printf("Error updating scoreboard %s", err)
			return levelHasBeenUp, starHasBeenUp, err
		}

		levelThresholdInfo, err := self.levelThresholdRepository.GetAll(ctx)
		if err != nil {
			fmt.Printf("Error fetching levelThresholdInfo")
			return levelHasBeenUp, starHasBeenUp, err
		}

		levelHasBeenUp, starHasBeenUp = account.Reputation.UpLevel(scoreboard, levelThresholdInfo)
	}

	account.AddAccumulateResponseTime(requestModel.ResponseTime)
	account.AddAnsweredQuiz(requestModel.QuizId) //TODO Review this approach

	err = self.accountRepository.Update(ctx, account)
	return
}

func (self *UpdateQuizAccountValuesUseCase) updateScoreboard(ctx context.Context, account entity.Account) (scoreboard entity.Scoreboard, err error) {
	scoreboardEndTimestamp := utils.GetEndOfCurrentWeek()
	scoreboard, err = self.scoreboardRepository.GetByIdentifier(ctx, account.NickName, scoreboardEndTimestamp)

	if err != nil {
		err = self.scoreboardRepository.Insert(ctx, entity.Scoreboard{
			Nickname:          account.NickName,
			HitsAmount:        1,
			ScoreEndTimestamp: scoreboardEndTimestamp,
		})
	} else {
		scoreboard.IncreaseHitsAmount()
		err = self.scoreboardRepository.Update(ctx, scoreboard)
	}

	return
}
