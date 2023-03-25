package db

import (
	"context"
	"fmt"
	"time"

	"github.com/Brolervi/Twitter_Chafa/models"
	"go.mongodb.org/mongo-driver/bson"
)

// CheckRelation, Checks if a relation exists
func CheckRelation(t models.Relation) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database(DBName)
	coll := db.Collection(RelationCollection)

	cond := bson.M{
		"userid":         t.UserID,
		"userrelationid": t.UserRelationID,
	}

	var res models.Relation

	err := coll.FindOne(ctx, cond).Decode(&res)
	if err != nil {
		fmt.Println(err.Error())
		return false, err
	}
	return true, nil
}
