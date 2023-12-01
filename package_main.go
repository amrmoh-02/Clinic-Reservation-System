package main

import (
	"context"
	//"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type Doctor struct {
	id       string   `bson:"id"`
	dname    string   `bson:"dname"`
	schedule []string `bson:"schedule"`
}

func main() {
	ctx := context.TODO()

	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017/")

	// Connect to MongoDB
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			log.Fatal(err)
		}
	}()

	// Access the "hospital" database and "doctor" collection
	coll := client.Database("hospital").Collection("doctor")

	// Create a Doctor instance with data
	doctor := Doctor{
		id:       "5",
		dname:    "Dr.ahmed",
		schedule: []string{"Monday", "Wednesday", "Friday"},
	}

	// Insert the Doctor instance into the "doctor" collection
	_, err = coll.InsertOne(ctx, doctor)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Doctor added successfully")
}
