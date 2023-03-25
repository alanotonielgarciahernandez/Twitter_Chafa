package db

import (
	"context"
	"time"

	"github.com/Brolervi/Twitter_Chafa/models"
	"go.mongodb.org/mongo-driver/bson"
)

// CheckUserAlreadyExists, Gets an email and asks the DB if the email is already in use
func CheckUserAlreadyExists(email string) (models.User, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database(DBName)
	coll := db.Collection(UserCollection)

	cond := bson.M{"email": email}

	var res models.User

	findMail := coll.FindOne(ctx, cond).Decode(&res)
	id := res.ID.Hex()

	if findMail != nil {
		return res, false, id
	}
	return res, true, id
}
