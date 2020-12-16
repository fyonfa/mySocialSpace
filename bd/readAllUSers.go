package bd

import (
	"context"
	"fmt"
	"time"

	"github.com/fyonfa/mySocialSpace/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ReadAllUsers(ID string, page int64, search string, typ string) ([]*models.User, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second) //dont want to be longer than 15s
	defer cancel()                                                           //it cancels the timeout given from cancel above

	db := MongoCN.Database("mySocial")
	col := db.Collection("users")

	var results []*models.User
	//the order here is important
	findOptions := options.Find()
	findOptions.SetSkip((page - 1) * 20)
	findOptions.SetLimit(20)

	query := bson.M{
		"name": bson.M{"$regex": `(?i)` + search}, //type regex, regular expresion, no case sensitive (?i)
	}
	//cursor
	cur, err := col.Find(ctx, query, findOptions)
	if err != nil {
		fmt.Println(err.Error())
		return results, false
	}

	var found, include bool
	//reading in every field
	for cur.Next(ctx) {
		var s models.User
		err := cur.Decode(&s)
		if err != nil {
			fmt.Println(err.Error())
			return results, false
		}
		var r models.Relation
		r.UserID = ID
		r.UserRelationID = s.ID.Hex()

		include = false

		found, err = ConsultRelation(r)
		if typ == "new" && found == false {
			include = true
		}
		if typ == "follow" && found == true {
			include = true
		}
		if r.UserRelationID == ID {
			include = false
		}
		if include == true {
			s.Password = ""
			s.Biography = ""
			s.Website = ""
			s.Location = ""
			s.Banner = ""
			s.Email = ""
			results = append(results, &s)
		}

	}
	err = cur.Err()
	if err != nil {
		fmt.Println(err.Error())
		return results, false
	}
	_ = cur.Close(ctx)
	return results, true
}
