package main

import (
	"github.com/gin-gonic/contrib/cors"
	"github.com/gin-gonic/gin"
	config "github.com/sabidos/configuration"

	_accountUsecase "github.com/sabidos/core/usecase/AccountUseCase"
	_avatarUsecase "github.com/sabidos/core/usecase/AvatarUseCase"
	_rankingUsecase "github.com/sabidos/core/usecase/RankingUseCase"

	_dataprovider "github.com/sabidos/dataprovider"
	_entrypoint "github.com/sabidos/entrypoint"
)

func main() {
	r := gin.Default()
	r.Use(cors.Default())
	db := config.ConnectToDB()
	api := r.Group("/v1")

	avatarDataProvider := _dataprovider.NewAvatarDataProvider(db)
	avatarUseCase := _avatarUsecase.NewObtainAvatarUsecase(avatarDataProvider)
	_entrypoint.NewAvatarEntrypointHandler(api, avatarUseCase)

	rankingDataProvider := _dataprovider.NewRankingDataProvider(db)
	rankingUseCase := _rankingUsecase.NewRankingUsecase(rankingDataProvider)
	_entrypoint.NewRankingEntrypointHandler(api, rankingUseCase)

	accountDataProvider := _dataprovider.NewAccountDataProvider(db)
	obtainAccountUseCase := _accountUsecase.NewObtainAccountUsecase(accountDataProvider)
	insertAccountUseCase := _accountUsecase.NewInsertAccountUsecase(accountDataProvider, avatarDataProvider)

	_entrypoint.NewAccountEntrypointHandler(api, obtainAccountUseCase, insertAccountUseCase)

	config.SetupModels(accountDataProvider, avatarDataProvider, db)

	r.Run()
}
