package db

import (
	"context"
	"time"

	"github.com/Brolervi/Twitter_Chafa/models"
)

func DeleteRelation(t models.Relation) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("Twitter_Chafa")
	coll := db.Collection("Relations")

	_, err := coll.DeleteOne(ctx, t)
	if err != nil {
		return false, err
	}
	return true, nil
}
