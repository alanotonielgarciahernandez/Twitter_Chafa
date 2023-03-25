package db

import (
	"context"
	"time"

	"github.com/Brolervi/Twitter_Chafa/models"
)

// InsertRelation, Inserts a new relation to the DB
func InsertRelation(t models.Relation) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database(DBName)
	coll := db.Collection(RelationCollection)

	_, err := coll.InsertOne(ctx, t)
	if err != nil {
		return false, err
	}

	return true, nil
}
