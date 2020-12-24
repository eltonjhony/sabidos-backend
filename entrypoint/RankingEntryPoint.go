package entrypoint

import (
	"github.com/gin-gonic/gin"

	"github.com/sabidos/core/usecase/RankingUseCase"
)

type RankingEntrypointHandler struct {
	ObtainRanking RankingUseCase.RankingUseCaseProtocol
}

func NewRankingEntrypointHandler(r *gin.RouterGroup, us RankingUseCase.RankingUseCaseProtocol) {
	handler := &RankingEntrypointHandler{
		ObtainRanking: us,
	}
	r.GET("/ranking", handler.FindRankings)
}

func (rankingEntrypointHandler *RankingEntrypointHandler) FindRankings(c *gin.Context) {
	rankings, _ := rankingEntrypointHandler.ObtainRanking.Fetch(c.Request.Context())
	c.JSON(200, gin.H{"data": rankings})
}
