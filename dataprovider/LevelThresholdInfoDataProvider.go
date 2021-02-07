package dataprovider

import (
	"context"
	"fmt"
	"log"

	"github.com/sabidos/core/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type LevelThresholdInfoDataProvider struct {
	Conn *mongo.Client
}

func NewLevelThresholdInfoDataProvider(Conn *mongo.Client) entity.LevelThresholdInfoDataProvider {
	return &LevelThresholdInfoDataProvider{Conn}
}

func (provider *LevelThresholdInfoDataProvider) GetAll(ctx context.Context) (res []entity.LevelThresholdInfo, err error) {
	collection := provider.Conn.Database("sabidos").Collection("levelThresholdInfo")

	bfilter := bson.M{}

	cur, err := collection.Find(context.TODO(), bfilter)
	if err != nil {
		log.Fatal("Error on Finding all the documents", err)
	}
	for cur.Next(context.TODO()) {
		var levelThresholdInfo entity.LevelThresholdInfo
		err = cur.Decode(&levelThresholdInfo)
		if err != nil {
			log.Fatal("Error on Decoding the document", err)
		}
		res = append(res, levelThresholdInfo)
		fmt.Printf("Found LevelThresholdInfo: %+v\n", levelThresholdInfo)

	}
	return res, err
}

func (provider *LevelThresholdInfoDataProvider) Insert(ctx context.Context, levelThresholdInfo entity.LevelThresholdInfo) (err error) {
	accountsCollection := provider.Conn.Database("sabidos").Collection("levelThresholdInfo")

	fmt.Printf("\nTrying to insert: %+v\n", levelThresholdInfo)

	result, err := accountsCollection.InsertOne(ctx, levelThresholdInfo)
	if err != nil {
		log.Panic("Error on Decoding the document", err)
		return err
	}
	fmt.Printf("\nInserted %v document into levelThresholdInfo collection!\n", result.InsertedID)
	fmt.Printf("\nInserted: %+v\n", levelThresholdInfo)

	return
}
