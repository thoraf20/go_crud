package database

import (
	"context"
	"fmt"
	"log"
	"time"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DBInstance() *mongo.Client {
	MongoDB := "mongodb+srv://physicist1:<physicist1>@cluster0.uvzxt.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"
	fmt.Print(MongoDB)

	client, err := mongo.NewClient(options.Client().ApplyURI(MongoDB))

	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.BackGround(), 10*time.Second)

	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.PrintLln("Connected to mongodb")

	return client
}

var Client *mongo.Client = DBInstance()

func OpenCollection(client *mongo.Client, collectionName string) {
	var collection *mongo.Client = client.Database("restaurant").Collection(collectionName)

	return collection
}
