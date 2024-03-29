package mongoclient

import (
	"context"
	"go-blog-service/utils"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var errorChecker utils.ErrorChecker

var Client *mongo.Client
var BlogCollection *mongo.Collection

func StartMongoClient() {
	log.Println("Connecting to MongoDB ...")
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	errorChecker.HasError(err).Fatal("An error occured when connecting to Mongo")
	connectionErr := client.Connect(context.TODO())
	errorChecker.HasError(connectionErr).Fatal("An error occured when connecting to Mongo")
	log.Println("Connected to MongoDB !")

	Client = client
	BlogCollection = client.Database("blog").Collection("blog")
}

//Run this func when program cancelled
func CloseConnection() {
	err := Client.Disconnect(context.TODO())
	errorChecker.HasError(err).Fatal("An error occured when disconnecting from Mongo")
}
