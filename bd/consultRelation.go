package bd

import (
	"context"
	"fmt"
	"github.com/fyonfa/mySocialSpace/models"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

//ConsultRelation consults the relation between 2 users
func ConsultRelation(t models.Relation) (bool, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second) //dont want to be longer than 15s
	defer cancel()                                                           //it cancels the timeout given from cancel above

	db := MongoCN.Database("mySocial")
	col := db.Collection("relation")

	condition := bson.M{
		"userid":         t.UserID,
		"userrelationid": t.UserRelationID,
	}
	var result models.Relation
	fmt.Println(result)
	err := col.FindOne(ctx, condition).Decode(&result)
	if err != nil {
		fmt.Println(err.Error())
		return false, err
	}
	return true, nil
}
