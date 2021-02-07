package entrypoint

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sabidos/core/usecase/QuizUseCase"
	"github.com/sabidos/entrypoint/model"
)

type QuizEntrypointHandler struct {
	ObtainQuiz        QuizUseCase.ObtainQuizUseCaseProtocol
	UpdateQuizAccount QuizUseCase.UpdateQuizAccountValuesUseCaseProtocol
}

func NewQuizRoundEntrypointHandler(r *gin.RouterGroup, obtainQuiz QuizUseCase.ObtainQuizUseCaseProtocol,
	updateQuizAccount QuizUseCase.UpdateQuizAccountValuesUseCaseProtocol) {
	handler := &QuizEntrypointHandler{
		ObtainQuiz:        obtainQuiz,
		UpdateQuizAccount: updateQuizAccount,
	}

	r.GET("/quiz/round/:nickname", handler.FindNextRound)
	r.POST("/quiz", handler.UpdateQuizValues)
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

func (quizEntrypointHandler *QuizEntrypointHandler) UpdateQuizValues(c *gin.Context) {

	var model model.PostQuizModel

	if err := c.ShouldBindJSON(&model); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Println(model)

	levelHasBeenUp, starHasBeenUp, err := quizEntrypointHandler.UpdateQuizAccount.UpdateQuizAccountValues(c, model)
	if err != nil {
		fmt.Println("Error Updating Quiz Account Values", err)
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"levelHasBeenUp": levelHasBeenUp,
		"starHasBeenUp":  starHasBeenUp,
	})
}
