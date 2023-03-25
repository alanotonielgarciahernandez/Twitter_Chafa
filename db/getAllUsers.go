package db

import (
	"context"
	"fmt"
	"time"

	"github.com/Brolervi/Twitter_Chafa/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// GetAllUsers, Gets all registered users from the DB
func GetAllUsers(ID string, page int64, search string, class string) ([]*models.User, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database(DBName)
	coll := db.Collection(UserCollection)

	var res []*models.User

	findOptions := options.Find()
	findOptions.SetSkip((page - 1) * 20)
	findOptions.SetLimit(20)

	query := bson.M{
		"name": bson.M{"$regex": `(?i)` + search},
	}

	cur, err := coll.Find(ctx, query, findOptions)
	if err != nil {
		fmt.Println(err.Error())
		return res, false
	}

	var found, include bool

	for cur.Next(ctx) {
		var s models.User
		err := cur.Decode(&s)
		if err != nil {
			fmt.Println(err.Error())
			return res, false
		}

		var r models.Relation
		r.UserID = ID
		r.UserRelationID = s.ID.Hex()

		include = false

		found, _ = CheckRelation(r)

		if class == "new" && !found {
			include = true
		}
		if class == "follow" && found {
			include = true
		}

		if r.UserRelationID == ID {
			include = false
		}

		if include {
			s.Password = ""
			s.Biography = ""
			s.Banner = ""
			s.Email = ""
			s.Location = ""
			s.Website = ""

			res = append(res, &s)
		}
	}

	err = cur.Err()
	if err != nil {
		fmt.Println(err.Error())
		return res, false
	}
	cur.Close(ctx)
	return res, true
}
