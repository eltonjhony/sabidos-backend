package model

import (
	"github.com/sabidos/core/entity"
)

type PostQuizModel struct {
	QuizId       string             `json:"quizId"`
	NickName     string             `json:"nickname"`
	ResponseTime int                `json:"responseTime"`
	Alternative  entity.Alternative `json:"alternative"`
}
