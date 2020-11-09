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

func (provider *AccountDataProvider) Get(ctx context.Context, filter bson.M) (res entity.Account, err error) {
	var account entity.Account
	fmt.Printf("Searching for account")

	collection := provider.Conn.Database("sabidos").Collection("accounts")

	if err = collection.FindOne(ctx, filter).Decode(&account); err != nil {
		log.Printf("Document with param  %s not found", filter)
		return account, err
	}
	fmt.Printf("Account found")

	return account, err
}

func (provider *AccountDataProvider) Insert(ctx context.Context, acc entity.Account) (err error) {
	accountsCollection := provider.Conn.Database("sabidos").Collection("accounts")

	result, err := accountsCollection.InsertOne(ctx, acc)
	if err != nil {
		log.Panic("Error on Decoding the document", err)
		return err
	}
	fmt.Printf("Inserted %v document into account collection!\n", result.InsertedID)

	return
}
