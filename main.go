package main

import (
	"github.com/gin-gonic/contrib/cors"
	"github.com/gin-gonic/gin"

	config "github.com/sabidos/configuration"

	_accountUsecase "github.com/sabidos/core/usecase/AccountUseCase"
	_avatarUsecase "github.com/sabidos/core/usecase/AvatarUseCase"
	_categoryUsecase "github.com/sabidos/core/usecase/CategoryUseCase"
	_quizUseCase "github.com/sabidos/core/usecase/QuizUseCase"
	_rankingUsecase "github.com/sabidos/core/usecase/RankingUseCase"

	_dataprovider "github.com/sabidos/dataprovider"

	_entrypoint "github.com/sabidos/entrypoint"
)

func main() {
	r := gin.Default()
	r.Use(cors.Default())
	db := config.ConnectToDB()
	api := r.Group("/v1")

	// Constructing avatars Resource
	avatarDataProvider := _dataprovider.NewAvatarDataProvider(db)
	avatarUseCase := _avatarUsecase.NewObtainAvatarUsecase(avatarDataProvider)
	_entrypoint.NewAvatarEntrypointHandler(api, avatarUseCase)

	// Constructing categories Resource
	categoryDataProvider := _dataprovider.NewCategoryDataProvider(db)
	categoryUseCase := _categoryUsecase.NewObtainCategoryUsecase(categoryDataProvider)
	_entrypoint.NewCategoryEntrypointHandler(api, categoryUseCase)

	// Constructing ranking Resource
	rankingDataProvider := _dataprovider.NewRankingDataProvider(db)
	rankingUseCase := _rankingUsecase.NewRankingUsecase(rankingDataProvider)
	_entrypoint.NewRankingEntrypointHandler(api, rankingUseCase)

	// Constructing accounts Resource
	accountDataProvider := _dataprovider.NewAccountDataProvider(db)
	obtainAccountUseCase := _accountUsecase.NewObtainAccountUsecase(accountDataProvider)
	insertAccountUseCase := _accountUsecase.NewInsertAccountUsecase(accountDataProvider, avatarDataProvider)
	validateAccountUseCase := _accountUsecase.NewValidateAccountUsecase(accountDataProvider)
	updateAccountUseCase := _accountUsecase.NewUpdateAccountUsecase(accountDataProvider)
	_entrypoint.NewAccountEntrypointHandler(api, obtainAccountUseCase, insertAccountUseCase, validateAccountUseCase, updateAccountUseCase)

	// Constructing quiz round Resource
	quizDataProvider := _dataprovider.NewQuizDataProvider(db)
	obtainQuizUseCase := _quizUseCase.NewObtainQuizUsecase(quizDataProvider)
	_entrypoint.NewQuizRoundEntrypointHandler(api, obtainQuizUseCase)

	// Init default database values
	config.SetupModels(accountDataProvider, avatarDataProvider, categoryDataProvider, quizDataProvider, db)

	r.Run()
}
