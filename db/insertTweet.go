package db

import (
	"context"
	"time"

	"github.com/Brolervi/Twitter_Chafa/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// InsertTweet, Inserts a new Tweet to the DB
func InsertTweet(t models.RecordTweet) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database(DBName)
	coll := db.Collection(TweetCollection)

	reg := bson.M{
		"userid":  t.UserID,
		"message": t.Message,
		"date":    t.Date,
	}
	res, err := coll.InsertOne(ctx, reg)
	if err != nil {
		return "", false, err
	}

	objID, _ := res.InsertedID.(primitive.ObjectID)

	return objID.String(), true, nil
}
