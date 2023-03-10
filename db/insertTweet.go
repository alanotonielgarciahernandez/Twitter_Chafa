package db

import (
	"context"
	"time"

	"github.com/Brolervi/Twitter_Chafa/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// InsertTweet, record a Tweet to the db
func InsertTweet(t models.RecordTweet) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("Twitter_Chafa")
	coll := db.Collection("Tweets")

	reg := bson.M{
		"userid":  t.UserID,
		"message": t.Message,
		"date":    t.Date,
	}
	result, err := coll.InsertOne(ctx, reg)
	if err != nil {
		return "", false, err
	}

	objID, _ := result.InsertedID.(primitive.ObjectID)

	return objID.String(), true, nil
}
