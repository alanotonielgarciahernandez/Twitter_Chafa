package db

import (
	"context"
	"time"

	"github.com/Brolervi/Twitter_Chafa/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ModifyLog, modifies user profile
func ModifyLog(u models.User, ID string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("Twitter_Chafa")
	coll := db.Collection("Users")

	reg := make(map[string]interface{})

	if len(u.Name) > 0 {
		reg["name"] = u.Name
	}
	if len(u.LName) > 0 {
		reg["lname"] = u.LName
	}
	reg["birthdate"] = u.BirthDate
	if len(u.Avatar) > 0 {
		reg["avatar"] = u.Avatar
	}
	if len(u.Banner) > 0 {
		reg["banner"] = u.Banner
	}
	if len(u.Biography) > 0 {
		reg["biography"] = u.Biography
	}
	if len(u.Location) > 0 {
		reg["location"] = u.Location
	}
	if len(u.Website) > 0 {
		reg["website"] = u.Website
	}

	updString := bson.M{
		"$set": reg,
	}

	objID, _ := primitive.ObjectIDFromHex(ID)
	filter := bson.M{"_id": bson.M{"$eq": objID}}

	_, err := coll.UpdateOne(ctx, filter, updString)
	if err != nil {
		return false, err
	}
	return true, nil
}
