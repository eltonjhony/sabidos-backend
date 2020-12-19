package entrypoint

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/sabidos/core/entity"
)

type QuizEntrypointHandler struct {
	ObtainQuiz entity.ObtainQuizUseCase
}

func NewQuizRoundEntrypointHandler(r *gin.RouterGroup, obtainQuiz entity.ObtainQuizUseCase) {
	handler := &QuizEntrypointHandler{
		ObtainQuiz: obtainQuiz,
	}

	r.GET("/quiz/round/:nickname", handler.FindNextRound)
}

func (quizEntrypointHandler *QuizEntrypointHandler) FindNextRound(c *gin.Context) {
	nickname := c.Param("nickname")
	categoryId := c.Query("categoryId")
	quizRound, err := quizEntrypointHandler.ObtainQuiz.ObtainQuizRoundFor(c.Request.Context(), nickname, categoryId)

	if err != nil {
		fmt.Println("Error Obtaining Quiz Round", err)
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}

	c.JSON(200, gin.H{"questions": quizRound})
}
