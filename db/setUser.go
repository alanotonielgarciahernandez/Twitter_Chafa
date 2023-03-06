package db

import (
	"context"
	"time"

	"github.com/Brolervi/Twitter_Chafa/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// SetUser, last step to insert user to DB
func SetUser(u models.User) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("Twitter_Chafa")
	col := db.Collection("Users")

	u.Password, _ = EncryptPassword(u.Password)

	result, err := col.InsertOne(ctx, u)
	if err != nil {
		return "", false, err
	}

	ObjID, _ := result.InsertedID.(primitive.ObjectID)
	return ObjID.String(), true, nil
}
