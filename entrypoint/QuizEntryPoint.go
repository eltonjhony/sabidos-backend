package entrypoint

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sabidos/core/entity"
	"github.com/sabidos/entrypoint/model"
)

type QuizEntrypointHandler struct {
	ObtainQuiz      entity.ObtainQuizUseCase
	UpdateQuizRound entity.UpdateQuizRoundUseCase
}

func NewQuizRoundEntrypointHandler(r *gin.RouterGroup, obtainQuiz entity.ObtainQuizUseCase, updateQuizRound entity.UpdateQuizRoundUseCase) {
	handler := &QuizEntrypointHandler{
		ObtainQuiz:      obtainQuiz,
		UpdateQuizRound: updateQuizRound,
	}

	r.GET("/quiz/round/:nickname", handler.FindNextRound)
	r.POST("/quiz/round", handler.UpdateQuizRoundValues)
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

func (quizEntrypointHandler *QuizEntrypointHandler) UpdateQuizRoundValues(c *gin.Context) {

	var model model.PostRoundModel

	if err := c.ShouldBindJSON(&model); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Println(model)

	err := quizEntrypointHandler.UpdateQuizRound.UpdateQuizRoundValues(c, model.NickName, model.AccumulateXp)
	if err != nil {
		fmt.Println("Error Updating Quiz Round Values", err)
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}

	c.JSON(200, gin.H{})
}
