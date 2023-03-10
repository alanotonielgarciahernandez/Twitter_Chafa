package models

//Tweet, tweet model
type Tweet struct {
	Message string `bson:"message" json:"message"`
}
