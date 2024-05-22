package database

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client

func GetCollection(name string) *mongo.Collection{
	return MongoClient.Database("test").Collection(name)
}

func StartMongoDB() error {
	uri := os.Getenv("MONGODB_URI")

	if uri == ""{
		log.Fatal("You must set ypu 'MONGODB_URI' environment variable.")
	}
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil{
		panic(err)
	}
	MongoClient = client

	return nil
}


func CloseMongoDB(){
	err := MongoClient.Disconnect(context.TODO())
	if err != nil{
		panic(err)
	}
}