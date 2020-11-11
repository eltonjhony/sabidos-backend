package config

import (
	"context"
	"log"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/sabidos/core/entity"
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

func SetupModels(dp entity.AccountDataProvider, conn *mongo.Client) {
	quickstartDatabase := conn.Database("sabidos")
	quickstartDatabase.Collection("accounts")

	account := entity.Account{"yiXtigKxtEVKl5mBh4qB7ZKumBs1", "Hulk", "Smash", entity.Avatar{1, ""}, entity.Reputation{"5", "10"}, "100", "100", "email", true, "tel"}

	dp.Insert(context.Background(), account)
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
