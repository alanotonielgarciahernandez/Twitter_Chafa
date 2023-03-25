package db

import (
	"context"
	"time"

	"github.com/Brolervi/Twitter_Chafa/models"
)

// DeleteRelation, Deletes the given relation from de DB
func DeleteRelation(t models.Relation) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database(DBName)
	coll := db.Collection(RelationCollection)

	_, err := coll.DeleteOne(ctx, t)
	if err != nil {
		return false, err
	}
	return true, nil
}
