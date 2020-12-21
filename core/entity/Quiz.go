package entity

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type QuizDataProvider interface {
	GetByCategory(ctx context.Context, categoryId int, limit int64) (res []Quiz, err error)
	Insert(ctx context.Context, acc Quiz) (err error)
}

type ObtainQuizUseCase interface {
	ObtainQuizRoundFor(ctx context.Context, nickname string, categoryId string) (res []Quiz, err error)
}

type Quiz struct {
	Id                 *primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	ImageUrl           string              `json:"imageUrl"`
	Description        string              `json:"description"`
	QuizLimitInSeconds int                 `json:"quizLimitInSeconds"`
	Category           Category            `json:"category"`
	Alternatives       []Alternative       `json:"alternatives"`
	Explanation        Explanation         `json:"explanation"`
}

type Alternative struct {
	Description        string `json:"description"`
	IsCorrect          bool   `json:"isCorrect"`
	PercentageAnswered int    `json:"percentageAnswered"`
}

type Explanation struct {
	Description string `json:"description"`
	Resource    string `json:"resource"`
}

func (quiz *Quiz) SetCategory(category Category) {
	quiz.Category = category
}

func (quiz *Quiz) AddAlternative(alternative Alternative) {
	quiz.Alternatives = append(quiz.Alternatives, alternative)
}

func (quiz *Quiz) SetExplanation(explanation Explanation) {
	quiz.Explanation = explanation
}
