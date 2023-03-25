package db

import (
	"context"
	"log"
	"time"

	"github.com/Brolervi/Twitter_Chafa/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// GetTweets, Gets a user's Tweets from the DB
func GetTweets(ID string, pag int64) ([]*models.ReturnTweet, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database(DBName)
	coll := db.Collection(TweetCollection)

	var res []*models.ReturnTweet

	cond := bson.M{
		"userid": ID,
	}

	opt := options.Find()
	opt.SetLimit(20)
	opt.SetSort(bson.D{{Key: "date", Value: -1}})
	opt.SetSkip((pag - 1) * 20)

	cursor, err := coll.Find(ctx, cond, opt)
	if err != nil {
		log.Fatal(err.Error())
		return res, false
	}

	for cursor.Next(context.TODO()) {
		var reg models.ReturnTweet
		err := cursor.Decode(&reg)
		if err != nil {
			return res, false
		}
		res = append(res, &reg)
	}

	return res, true
}
