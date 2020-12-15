package bd

import (
	"context"
	"github.com/fyonfa/mySocialSpace/models"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)
//CheckUserExists receive an email of parameter and check if is in the DB already
func CheckUserExists(email string)(models.User,bool, string){
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("mySocial")
	col := db.Collection("users")

	//conditions in mongo
	condition := bson.M{"email":email}

	var result models.User

	err := col.FindOne(ctx, condition).Decode(&result)
	ID := result.ID.Hex()
	if err != nil{
		return result,false,ID
	}
	return result,true, ID
}
