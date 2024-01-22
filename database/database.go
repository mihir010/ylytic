package database

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const URI string = "mongodb+srv://kumarmihir02:mihir@gofirst.sj5svf3.mongodb.net/?retryWrites=true&w=majority"
const db string = "test"
const coll string = "string"

func Collection() *mongo.Collection {
	opts := options.Client().ApplyURI(URI)

	client, err := mongo.Connect(context.TODO(), opts)

	if err != nil {
		log.Fatal(err.Error())
	}

	collec := client.Database(db).Collection(coll)

	return collec
}
