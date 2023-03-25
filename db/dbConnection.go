package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoCN, Connection object to DB
var MongoCN = DBconnect()
var clientOptions = options.Client().ApplyURI("mongodb+srv://admon:ad03m07in@cluster0.xpx9x7e.mongodb.net/test") // MongoDB Cluster direction

// DBName, Database Name used in al DB folder endpoints
var DBName = "Twitter_Chafa" // You can change these names with your DB collection names

var RelationCollection = "Relations"
var TweetCollection = "Tweets"
var UserCollection = "Users"

// DBconnect, Connects to DB
func DBconnect() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	log.Println("DB connection established")
	return client
}

// CheckDBConnection, Obtains Ping from DB
func CheckDBConnection() int {
	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1
}
