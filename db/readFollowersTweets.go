package db

import (
	"context"
	"time"

	"github.com/Brolervi/Twitter_Chafa/models"
	"go.mongodb.org/mongo-driver/bson"
)

// ReadFollowersTweets, Gets all Tweets from the given user's followers
func ReadFollowersTweets(ID string, page int) ([]models.FollowersTweets, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database(DBName)
	coll := db.Collection(RelationCollection)

	skip := (page - 1) * 20

	cond := make([]bson.M, 0)
	cond = append(cond, bson.M{"$match": bson.M{"userid": ID}})
	cond = append(cond, bson.M{
		"$lookup": bson.M{
			"from":         "Tweets",
			"localField":   "userrelationid",
			"foreignField": "userid",
			"as":           "tweet",
		}})
	cond = append(cond, bson.M{"$unwind": "$tweet"})
	cond = append(cond, bson.M{"$sort": bson.M{"date": -1}})
	cond = append(cond, bson.M{"$skip": skip})
	cond = append(cond, bson.M{"$limit": 20})

	var res []models.FollowersTweets

	cursor, err := coll.Aggregate(ctx, cond)
	if err != nil {
		return res, false
	}

	err = cursor.All(ctx, &res)
	if err != nil {
		return res, false
	}
	return res, true
}
