package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// DeleteTweet, Deletes the given Tweet from DB
func DeleteTweet(ID string, UserID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database(DBName)
	coll := db.Collection(TweetCollection)

	objID, _ := primitive.ObjectIDFromHex(ID)

	cond := bson.M{
		"_id":    objID,
		"userid": UserID,
	}

	_, err := coll.DeleteOne(ctx, cond)
	return err
}
