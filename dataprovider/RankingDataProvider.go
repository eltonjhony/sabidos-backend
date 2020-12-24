package dataprovider

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/sabidos/core/entity"
	"go.mongodb.org/mongo-driver/mongo"
)

type RankingDataProvider struct {
	Conn *mongo.Client
}

func NewRankingDataProvider(Conn *mongo.Client) entity.RankingDataProvider {
	return &RankingDataProvider{Conn}
}

func (provider *RankingDataProvider) Fetch(ctx context.Context) (res []entity.Ranking, err error) {
	var rankingList []entity.Ranking
	bfilter := bson.M{"name": "Test"}
	collection := provider.Conn.Database("sabidos").Collection("rankings")
	cur, err := collection.Find(context.TODO(), bfilter)
	if err != nil {
		log.Fatal("Error on Finding all the documents", err)
	}

	for cur.Next(context.TODO()) {
		var ranking entity.Ranking
		err = cur.Decode(&ranking)
		if err != nil {
			log.Fatal("Error on Decoding the document", err)
		}
		rankingList = append(rankingList, ranking)
	}
	return rankingList, nil
}
