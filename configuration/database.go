package config

import (
	"context"
	"log"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/sabidos/core/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var db *mongo.Client
var port string

var collection *mongo.Collection
var ctx = context.TODO()

func ConnectToDB() *mongo.Client {
	clientOptions := options.Client().ApplyURI("mongodb://db:27017")
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Connect(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatal("Couldn't connect to the database", err)
	} else {
		log.Println("Connected to MongoDB!")
	}

	return client
}

func SetupModels(ac entity.AccountDataProvider, av entity.AvatarDataProvider, conn *mongo.Client) {
	quickstartDatabase := conn.Database("sabidos")
	quickstartDatabase.Collection("accounts")

	//TODO clean db

	avatar1 := entity.Avatar{1, "https://res.cloudinary.com/ddb86uj5i/image/upload/v1604167774/sabidos/avatar/1.png"}
	avatar2 := entity.Avatar{2, "https://res.cloudinary.com/ddb86uj5i/image/upload/v1604167776/sabidos/avatar/2.png"}
	avatar3 := entity.Avatar{3, "https://res.cloudinary.com/ddb86uj5i/image/upload/v1604167776/sabidos/avatar/3.png"}
	avatar4 := entity.Avatar{4, "https://res.cloudinary.com/ddb86uj5i/image/upload/v1604167776/sabidos/avatar/4.png"}
	avatar5 := entity.Avatar{5, "https://res.cloudinary.com/ddb86uj5i/image/upload/v1604167775/sabidos/avatar/5.png"}
	avatar6 := entity.Avatar{6, "https://res.cloudinary.com/ddb86uj5i/image/upload/v1604167775/sabidos/avatar/6.png"}
	avatar7 := entity.Avatar{7, "https://res.cloudinary.com/ddb86uj5i/image/upload/v1604167775/sabidos/avatar/7.png"}
	avatar8 := entity.Avatar{8, "https://res.cloudinary.com/ddb86uj5i/image/upload/v1604167775/sabidos/avatar/8.png"}
	avatar9 := entity.Avatar{9, "https://res.cloudinary.com/ddb86uj5i/image/upload/v1604167775/sabidos/avatar/9.png"}
	avatar10 := entity.Avatar{10, "https://res.cloudinary.com/ddb86uj5i/image/upload/v1604167775/sabidos/avatar/10.png"}

	bfilter := bson.M{"$or": []bson.M{bson.M{"nickname": "smash"}, bson.M{"uid": ""}}}

	account, _ := ac.Get(context.Background(), bfilter)

	if len(account.NickName) == 0 {
		newAcc := entity.Account{"yiXtigKxtEVKl5mBh4qB7ZKumBs1", "Hulk", "Smash", entity.Avatar{1, ""}, entity.Reputation{5, 10}, 100, 100, "email", true, "tel"}
		ac.Insert(context.Background(), newAcc)

	}

	avatarFilter := bson.M{"id": 1}

	avatar, _ := av.FindOne(context.Background(), avatarFilter)

	if avatar.Id == 0 {

		av.Insert(context.Background(), avatar1)
		av.Insert(context.Background(), avatar2)
		av.Insert(context.Background(), avatar3)
		av.Insert(context.Background(), avatar4)
		av.Insert(context.Background(), avatar5)
		av.Insert(context.Background(), avatar6)
		av.Insert(context.Background(), avatar7)
		av.Insert(context.Background(), avatar8)
		av.Insert(context.Background(), avatar9)
		av.Insert(context.Background(), avatar10)
	}

}

func SetUpDBConnection(DB *mongo.Client) {
	db = DB
}

func GetDBConnection() *mongo.Client {
	return db
}

func SetPortConnection(Port string) {
	port = Port
}

func GetPortConnection() string {
	return port
}
