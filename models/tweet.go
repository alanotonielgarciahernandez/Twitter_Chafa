package models

//Tweet, Only Tweet message model
type Tweet struct {
	Message string `bson:"message" json:"message"`
}
