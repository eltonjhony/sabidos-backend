package dataprovider

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/sabidos/core/entity"
	"go.mongodb.org/mongo-driver/mongo"
)

type CategoryDataProvider struct {
	Conn *mongo.Client
}

func NewCategoryDataProvider(Conn *mongo.Client) entity.CategoryDataProvider {
	return &CategoryDataProvider{Conn}
}

func (provider *CategoryDataProvider) Get(ctx context.Context, filter bson.M) (res []entity.Category, err error) {
	var categories []entity.Category
	collection := provider.Conn.Database("sabidos").Collection("categories")
	cur, err := collection.Find(context.TODO(), filter)
	if err != nil {
		log.Fatal("Error on Finding all the documents", err)
	}
	for cur.Next(context.TODO()) {
		var category entity.Category
		err = cur.Decode(&category)
		if err != nil {
			log.Fatal("Error on Decoding the document", err)
		}
		categories = append(categories, category)
		fmt.Printf("Found category: %+v\n", category)

	}
	return categories, err
}

func (provider *CategoryDataProvider) Insert(ctx context.Context, acc entity.Category) (err error) {
	categoriesCollection := provider.Conn.Database("sabidos").Collection("categories")

	fmt.Printf("Trying to insert: %+v\n", acc)

	result, err := categoriesCollection.InsertOne(ctx, acc)
	if err != nil {
		log.Panic("Error on Decoding the document", err)
		return err
	}
	fmt.Printf("Inserted %v document into categories collection!\n", result.InsertedID)
	fmt.Printf("Inserted: %+v\n", acc)

	return
}

func (provider *CategoryDataProvider) FindOne(ctx context.Context, filter bson.M) (res entity.Category, err error) {
	var category entity.Category
	fmt.Printf("\nSearching for Category")

	collection := provider.Conn.Database("sabidos").Collection("categories")

	if err = collection.FindOne(ctx, filter).Decode(&category); err != nil {
		log.Printf("\nDocument with param  %s not found", filter)
		return category, err
	}

	fmt.Printf("\ncategory found")

	return category, err
}
