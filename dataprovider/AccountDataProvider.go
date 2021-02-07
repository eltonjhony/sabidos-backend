package dataprovider

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/sabidos/core/entity"
	"go.mongodb.org/mongo-driver/mongo"
)

type AccountDataProvider struct {
	Conn *mongo.Client
}

func NewAccountDataProvider(Conn *mongo.Client) entity.AccountDataProvider {
	return &AccountDataProvider{Conn}
}

func (provider *AccountDataProvider) GetByIdentifier(ctx context.Context, nickname string, uid string) (res entity.Account, err error) {
	var account entity.Account
	fmt.Printf("\nSearching for account")

	collection := provider.Conn.Database("sabidos").Collection("accounts")

	bfilter := bson.M{"$or": []bson.M{bson.M{"nickname": nickname}, bson.M{"uid": uid}}}

	if err = collection.FindOne(ctx, bfilter).Decode(&account); err != nil {
		log.Printf("\nDocument with param  %s not found", bfilter)
		return account, err
	}
	fmt.Printf("\nAccount found")
	log.Printf("\nDocument with param  %s found", bfilter)

	return account, err
}

func (provider *AccountDataProvider) GetByNickname(ctx context.Context, nickname string) (res entity.Account, err error) {
	var account entity.Account
	fmt.Printf("\nSearching for account")

	collection := provider.Conn.Database("sabidos").Collection("accounts")

	bfilter := bson.M{"nickname": nickname}

	if err = collection.FindOne(ctx, bfilter).Decode(&account); err != nil {
		log.Printf("\nDocument with param  %s not found", bfilter)
		return account, err
	}
	fmt.Printf("\nAccount found")
	log.Printf("\nDocument with param  %s found", bfilter)

	return account, err
}

func (provider *AccountDataProvider) GetByUid(ctx context.Context, uid string) (res entity.Account, err error) {
	var account entity.Account
	fmt.Printf("\nSearching for account")

	collection := provider.Conn.Database("sabidos").Collection("accounts")

	bfilter := bson.M{"uid": uid}

	if err = collection.FindOne(ctx, bfilter).Decode(&account); err != nil {
		log.Printf("\nDocument with param  %s not found", bfilter)
		return account, err
	}
	fmt.Printf("\nAccount found")
	log.Printf("\nDocument with param  %s found", bfilter)

	return account, err
}

func (provider *AccountDataProvider) Insert(ctx context.Context, acc entity.Account) (err error) {
	accountsCollection := provider.Conn.Database("sabidos").Collection("accounts")

	fmt.Printf("\nTrying to insert: %+v\n", acc)

	result, err := accountsCollection.InsertOne(ctx, acc)
	if err != nil {
		log.Panic("Error on Decoding the document", err)
		return err
	}
	fmt.Printf("\nInserted %v document into account collection!\n", result.InsertedID)
	fmt.Printf("\nInserted: %+v\n", acc)

	return
}

func (provider *AccountDataProvider) Update(ctx context.Context, acc entity.Account) (err error) {
	accountsCollection := provider.Conn.Database("sabidos").Collection("accounts")

	fmt.Printf("\nTrying to update: %+v\n", acc)

	bfilter := bson.M{"uid": bson.M{"$eq": acc.Uid}}

	update := bson.M{
		"$set": bson.M{
			"name":                   acc.Name,
			"email":                  acc.Email,
			"phone":                  acc.Phone,
			"isanonymous":            acc.IsAnonymous,
			"accumulatexp":           acc.AccumulateXp,
			"totalanswered":          acc.TotalAnswered,
			"totalhits":              acc.TotalHits,
			"accumulateresponsetime": acc.AccumulateResponseTime,
			"answeredquiz":           acc.AnsweredQuiz,
			"reputation":             acc.Reputation,
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
	fmt.Printf("\nUpdate: %+v\n", acc.NickName)

	return
}
