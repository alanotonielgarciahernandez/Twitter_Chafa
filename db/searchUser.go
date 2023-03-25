package db

import (
	"context"
	"fmt"
	"time"

	"github.com/Brolervi/Twitter_Chafa/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// SearchUser, Searches for a user in the DB and obtains its data
func SearchUser(ID string) (models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := MongoCN.Database(DBName)
	coll := db.Collection(UserCollection)

	var profile models.User
	objID, _ := primitive.ObjectIDFromHex(ID)

	cond := bson.M{
		"_id": objID,
	}

	err := coll.FindOne(ctx, cond).Decode(&profile)
	profile.Password = ""
	if err != nil {
		fmt.Println("Couldn't find user " + err.Error())
		return profile, err
	}
	return profile, nil
}
