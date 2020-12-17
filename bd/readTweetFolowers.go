package bd

import (
	"context"
	"github.com/fyonfa/mySocialSpace/models"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

//ReadTweetFollowers reads the tweets from my followers
func ReadTweetFollowers(ID string, page int) ([]models.ReturnTweetFollowers, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second) //dont want to be longer than 15s
	defer cancel()                                                           //it cancels the timeout given from cancel above

	db := MongoCN.Database("mySocial")
	col := db.Collection("relation")

	skip := (page - 1) * 20         //20 result only

	conditions := make([]bson.M, 0) //for agregate framework MongoDB
	conditions = append(conditions, bson.M{"$match": bson.M{"userid": ID}})
	//join two tables with mongo
	conditions = append(conditions, bson.M{
		"$lookup": bson.M{
			"from":         "tweet",
			"localField":   "userrelationid",
			"foreignField": "userid",
			"as":           "tweet",
		}})
	conditions = append(conditions, bson.M{"$unwind": "$tweet"})               //unwind:all documents come the same
	conditions = append(conditions, bson.M{"$sort": bson.M{"tweet.date": -1}}) //sort
	conditions = append(conditions, bson.M{"skip": skip})                      //first skip and then limit, if not works bad
	conditions = append(conditions, bson.M{"$limit": 20})
	//framework agregate
	cursor, err := col.Aggregate(ctx, conditions)
	var result []models.ReturnTweetFollowers
	err = cursor.All(ctx, &result)
	if err != nil {
		return result, false
	}
	return result, true
}
