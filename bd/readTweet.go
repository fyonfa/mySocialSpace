package bd

import (
	"context"
	"github.com/fyonfa/mySocialSpace/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

//ReadTweet reads all the tweets from a profile
func ReadTweet(ID string, page int64) ([]*models.ReturnTweet, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second) //dont want to be longer than 15s
	defer cancel()                                                           //it cancels the timeout given from cancel above

	db := MongoCN.Database("mySocial")
	col := db.Collection("tweet")

	var result []*models.ReturnTweet

	condition := bson.M{
		"userid": ID,
	}
	opciones := options.Find()                         //we are going to work with the mode options in find
	opciones.SetLimit(20)                              //20 first, tweets that we are going to sent as paramater Setskip
	opciones.SetSort(bson.D{{Key: "data", Value: -1}}) //tweets ordered, key marks the sorted, -1 descending order, last tweets firs
	opciones.SetSkip((page - 1) * 20)

	cursor, err := col.Find(ctx, condition, opciones)
	if err != nil {
		log.Fatal(err.Error())
		return result, false
	}
	//context new without limitations
	for cursor.Next(context.TODO()) {
		var register models.ReturnTweet
		err := cursor.Decode(&register)
		if err != nil {
			return result, false
		}
		result = append(result, &register)
	}
	return result, true
}
