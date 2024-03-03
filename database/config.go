package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DBInstance() *mongo.Client {
	MongoDB := os.Getenv("MONGODB_URI")
	fmt.Print(MongoDB)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(MongoDB))

	if err != nil {
		log.Fatal(err)
	}

	// err = client.Connect(ctx)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	fmt.Println("connection to database is successful")

	return client
}

var Client *mongo.Client = DBInstance()

func OpenCollection(client *mongo.Client, collectionName string) {
	// var collection = 
	client.Database("restaurant").Collection(collectionName)

	// return collection
}
