package main

import (
	"github.com/gin-gonic/contrib/cors"
	"github.com/gin-gonic/gin"
	config "github.com/sabidos/configuration"

	_usecase "github.com/sabidos/core/usecase/RankingUseCase"
	_dataprovider "github.com/sabidos/dataprovider"
	_entrypoint "github.com/sabidos/entrypoint"
)

func main() {
	config.SetupModels()
	r := gin.Default()
	r.Use(cors.Default())
	db := config.ConnectToDB()

	rankingDataProvider := _dataprovider.NewRankingDataProvider(db)

	rankingUseCase := _usecase.NewRankingUsecase(rankingDataProvider)

	api := r.Group("/v1")

	_entrypoint.NewRankingEntrypointHandler(api, rankingUseCase)

	r.Run()
}
