package dataprovider

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/sabidos/core/entity"
	"go.mongodb.org/mongo-driver/mongo"
)

type ScoreboardDataProvider struct {
	Conn *mongo.Client
}

func NewScoreboardDataProvider(Conn *mongo.Client) entity.ScoreboardDataProvider {
	return &ScoreboardDataProvider{Conn}
}

func (provider *ScoreboardDataProvider) GetByIdentifier(ctx context.Context, nickname string, scoreEndTimestamp int) (res entity.Scoreboard, err error) {
	var scoreboard entity.Scoreboard
	fmt.Printf("\nSearching for scoreboard")

	collection := provider.Conn.Database("sabidos").Collection("scoreboard")

	bfilter := bson.M{"$and": []bson.M{
		bson.M{"nickname": nickname},
		bson.M{"scoreendtimestamp": scoreEndTimestamp}}}

	if err = collection.FindOne(ctx, bfilter).Decode(&scoreboard); err != nil {
		log.Printf("\nDocument with param  %s not found", bfilter)
		return scoreboard, err
	}
	fmt.Printf("\nScoreboard found")
	log.Printf("\nDocument with param  %s found", bfilter)

	return scoreboard, err
}

func (provider *ScoreboardDataProvider) Insert(ctx context.Context, scoreboard entity.Scoreboard) (err error) {
	accountsCollection := provider.Conn.Database("sabidos").Collection("scoreboard")

	fmt.Printf("\nTrying to insert: %+v\n", scoreboard)

	result, err := accountsCollection.InsertOne(ctx, scoreboard)
	if err != nil {
		log.Panic("Error on Decoding the document", err)
		return err
	}
	fmt.Printf("\nInserted %v document into scoreboard collection!\n", result.InsertedID)
	fmt.Printf("\nInserted: %+v\n", scoreboard)

	return
}

func (provider *ScoreboardDataProvider) Update(ctx context.Context, scoreboard entity.Scoreboard) (err error) {
	accountsCollection := provider.Conn.Database("sabidos").Collection("scoreboard")

	fmt.Printf("\nTrying to update: %+v\n", scoreboard)

	bfilter := bson.M{"$and": []bson.M{bson.M{"nickname": scoreboard.Nickname}, bson.M{"scoreendtimestamp": scoreboard.ScoreEndTimestamp}}}

	update := bson.M{
		"$set": bson.M{
			"hitsamount": scoreboard.HitsAmount,
		},
	}

	result, err := accountsCollection.UpdateOne(
		ctx,
		bfilter,
		update,
	)

	if err != nil {
		log.Panic("Error on Updating the document", err)
		return err
	}

	fmt.Printf("\nUpdated %v document on account collection!\n", result.UpsertedID)
	fmt.Printf("\nUpdate: %+v\n", scoreboard.Nickname)

	return
}
