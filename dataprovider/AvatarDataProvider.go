package dataprovider

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/sabidos/core/entity"
	"go.mongodb.org/mongo-driver/mongo"
)

type AvatarDataProvider struct {
	Conn *mongo.Client
}

func NewAvatarDataProvider(Conn *mongo.Client) entity.AvatarDataProvider {
	return &AvatarDataProvider{Conn}
}

func (provider *AvatarDataProvider) GetAll(ctx context.Context) (res []entity.Avatar, err error) {
	var avatars []entity.Avatar
	collection := provider.Conn.Database("sabidos").Collection("avatars")

	bfilter := bson.M{}

	cur, err := collection.Find(context.TODO(), bfilter)
	if err != nil {
		log.Fatal("Error on Finding all the documents", err)
	}
	for cur.Next(context.TODO()) {
		var avatar entity.Avatar
		err = cur.Decode(&avatar)
		if err != nil {
			log.Fatal("Error on Decoding the document", err)
		}
		avatars = append(avatars, avatar)
		fmt.Printf("Found Avatar: %+v\n", avatar)

	}
	return avatars, err
}

func (provider *AvatarDataProvider) Count(ctx context.Context) (res int64, err error) {
	collection := provider.Conn.Database("sabidos").Collection("avatars")

	fmt.Printf("Starting count avatars")

	bfilter := bson.M{}

	itemCount, err := collection.CountDocuments(ctx, bfilter)
	if err != nil {
		log.Fatal("Error on Counting all documents", err)
	}

	fmt.Printf("Getting total count of avatars %d", itemCount)

	return itemCount, err
}

func (provider *AvatarDataProvider) Insert(ctx context.Context, acc entity.Avatar) (err error) {
	avatarsCollection := provider.Conn.Database("sabidos").Collection("avatars")

	fmt.Printf("Trying to insert: %+v\n", acc)

	result, err := avatarsCollection.InsertOne(ctx, acc)
	if err != nil {
		log.Panic("Error on Decoding the document", err)
		return err
	}
	fmt.Printf("Inserted %v document into avatar collection!\n", result.InsertedID)
	fmt.Printf("Inserted: %+v\n", acc)

	return
}

func (provider *AvatarDataProvider) FindById(ctx context.Context, id int) (res entity.Avatar, err error) {
	var avatar entity.Avatar
	fmt.Printf("\nSearching for Avatar")

	collection := provider.Conn.Database("sabidos").Collection("avatars")

	bfilter := bson.M{"id": id}

	if err = collection.FindOne(ctx, bfilter).Decode(&avatar); err != nil {
		log.Printf("\nDocument with param  %s not found", bfilter)
		return avatar, err
	}

	fmt.Printf("\nAvatar found")

	return avatar, err
}
