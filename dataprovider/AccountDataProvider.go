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

func (provider *AccountDataProvider) Fetch(ctx context.Context) (res []entity.Account, err error) {
	var accounts []entity.Account
	bfilter := bson.M{"id": "PUr2i6QQcMfKBGRL9sYyw8UJ0tv1"}
	collection := provider.Conn.Database("sabidos").Collection("account")
	cur, err := collection.Find(context.TODO(), bfilter)
	if err != nil {
		log.Fatal("Error on Finding all the documents", err)
	}

	for cur.Next(context.TODO()) {
		var account entity.Account
		err = cur.Decode(&account)
		if err != nil {
			log.Fatal("Error on Decoding the document", err)
		}
		accounts = append(accounts, account)
	}
	return accounts, nil
}

func (provider *AccountDataProvider) Insert(acc entity.Account, ctx context.Context) {
	accountsCollection := provider.Conn.Database("sabidos").Collection("account")
	result, err := accountsCollection.InsertOne(ctx, bson.D{
		{Key: "title", Value: "The Polyglot Developer Podcast"},
		{Key: "author", Value: "Nic Raboy"},
	})
	if err != nil {
		log.Fatal("Error on Decoding the document", err)
	}
	fmt.Printf("Inserted %v document into account collection!\n", result.InsertedID)
}
