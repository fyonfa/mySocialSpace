package bd

import (
	"context"
	"github.com/fyonfa/mySocialSpace/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)
//InsertRegister it is the final stop with the DB to insert the dat from the user
func InsertRegister(u models.User) (string, bool, error){
	//create a context
	ctx, cancel := context.WithTimeout(context.Background(),15*time.Second)//dont want to be longer than 15s
	defer cancel()//it cancels the timeout given from cancel above

	db := MongoCN.Database("mySocial")
	col := db.Collection("users")

	u.Password, _ = EncryptPassword(u.Password)

	result, err := col.InsertOne(ctx,u)
	if err != nil{
		return "", false, err//not id anf false, and error
	}
	ObjID,_ := result.InsertedID.(primitive.ObjectID)//this is the way to get after insert one
	return ObjID.String(), true, nil
 }
