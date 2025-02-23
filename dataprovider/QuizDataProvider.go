package dataprovider

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/sabidos/core/entity"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type QuizDataProvider struct {
	Conn *mongo.Client
}

func NewQuizDataProvider(Conn *mongo.Client) entity.QuizDataProvider {
	return &QuizDataProvider{Conn}
}

func (provider *QuizDataProvider) GetByParams(ctx context.Context, params entity.QuizParams) (res []entity.Quiz, err error) {
	var quizRound []entity.Quiz
	collection := provider.Conn.Database("sabidos").Collection("quiz")

	options := options.Find()

	// Limit the amount of documents being returned
	options.SetLimit(params.Limit)

	// Filter params
	bfilters := bson.M{}
	if params.CategoryId > 1 {
		fmt.Printf("Filter including category: %+d\n", params.CategoryId)
		bfilters = bson.M{"category.id": params.CategoryId}
	}

	cur, err := collection.Find(ctx, bfilters, options)

	if err != nil {
		log.Fatal("Error on Finding all the documents", err)
	}

	for cur.Next(context.TODO()) {
		var quiz entity.Quiz
		err = cur.Decode(&quiz)
		if err != nil {
			log.Fatal("Error on Decoding the document", err)
		}
		quizRound = append(quizRound, quiz)
		fmt.Printf("Found Quiz: %+v\n", quiz)
	}

	if quizRound == nil {
		fmt.Printf("Quiz Round did not Found results")
		return []entity.Quiz{}, err
	}

	return quizRound, err
}

func (provider *QuizDataProvider) Insert(ctx context.Context, acc entity.Quiz) (err error) {
	quizCollection := provider.Conn.Database("sabidos").Collection("quiz")

	fmt.Printf("Trying to insert: %+v\n", acc)

	result, err := quizCollection.InsertOne(ctx, acc)
	if err != nil {
		log.Panic("Error on Decoding the document", err)
		return err
	}
	fmt.Printf("Inserted %v document into quiz collection!\n", result.InsertedID)
	fmt.Printf("Inserted: %+v\n", acc)

	return
}
