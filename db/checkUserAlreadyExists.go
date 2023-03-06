package db

import (
	"context"
	"time"

	"github.com/Brolervi/Twitter_Chafa/models"
	"go.mongodb.org/mongo-driver/bson"
)

// CheckUserAlreadyExists, get an email and ask to the db if email is already in use
func CheckUserAlreadyExists(email string) (models.User, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("Twitter_Chafa")
	coll := db.Collection("users")

	cond := bson.M{"email": email}

	var res models.User

	err := coll.FindOne(ctx, cond).Decode(&res)
	id := res.ID.Hex()

	if err != nil {
		return res, false, id
	}
	return res, true, id
}
