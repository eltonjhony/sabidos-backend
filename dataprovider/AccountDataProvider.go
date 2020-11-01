package dataprovider

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"strconv"

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

func (provider *AccountDataProvider) Get(ctx context.Context, id string) (res entity.Account, err error) {
	var account entity.Account
	i, err := strconv.ParseInt(id, 10, 64)
	bfilter := bson.M{"id": i}

	collection := provider.Conn.Database("sabidos").Collection("accounts")

	if err = collection.FindOne(ctx, bfilter).Decode(&account); err != nil {
		log.Panic(err)
	}

	fmt.Println(account)

	return account, nil
}

func (provider *AccountDataProvider) Insert(ctx context.Context, acc entity.Account) (err error) {
	accountsCollection := provider.Conn.Database("sabidos").Collection("accounts")
	if acc.Id == 0 {
		acc.SetId(rand.Intn(100000))
	}
	result, err := accountsCollection.InsertOne(ctx, bson.D{
		{Key: "id", Value: acc.Id},
		{Key: "name", Value: acc.Name},
		{Key: "nickname", Value: acc.NickName},
	})
	if err != nil {
		log.Fatal("Error on Decoding the document", err)
		return err
	}
	fmt.Printf("Inserted %v document into account collection!\n", result.InsertedID)

	return
}
