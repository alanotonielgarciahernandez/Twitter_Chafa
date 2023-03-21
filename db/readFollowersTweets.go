package db

import (
	"context"
	"time"

	"github.com/Brolervi/Twitter_Chafa/models"
	"go.mongodb.org/mongo-driver/bson"
)

// ReadFollowersTweets, Read all followers Tweets of a user
func ReadFollowersTweets(ID string, page int) ([]models.FollowersTweets, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("Twitter_Chafa")
	coll := db.Collection("Relations")

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

	cursor, err := coll.Aggregate(ctx, cond)
	var res []models.FollowersTweets
	err = cursor.All(ctx, &res)
	if err != nil {
		return res, false
	}
	return res, true
}
