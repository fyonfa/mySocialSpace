package bd

import (
	"context"
	"github.com/fyonfa/mySocialSpace/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

//InsertTweet insert tweet in the data base
func InsertTweet(t models.RecordTweet) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second) //dont want to be longer than 15s
	defer cancel()                                                           //it cancels the timeout given from cancel above

	db := MongoCN.Database("mySocial")
	col := db.Collection("tweet")

	//convert into bson
	register := bson.M{
		"userid":  t.UserID,
		"message": t.Message,
		"date":    t.Date,
	}
	result, err := col.InsertOne(ctx, register)
	if err != nil {
		return "", false, err
	}
	objID, _ := result.InsertedID.(primitive.ObjectID)
	return objID.String(), true, nil
}
