package bd

import (
	"context"
	"time"

	"github.com/fyonfa/mySocialSpace/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//ModifyRegister allows to modify the user profile
func ModifyRegister(u models.User, ID string) (bool, error) {
	//create a context
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second) //dont want to be longer than 15s
	defer cancel()                                                           //it cancels the timeout given from cancel above

	db := MongoCN.Database("mySocial")
	col := db.Collection("users")

	register := make(map[string]interface{})
	if len(u.Name) > 0 {
		register["name"] = u.Name
	}
	if len(u.Surname) > 0 {
		register["surname"] = u.Surname
	}
	register["birthDate"] = u.BirthDate
	if len(u.Avatar) > 0 {
		register["avatar"] = u.Avatar
	}
	if len(u.Banner) > 0 {
		register["banner"] = u.Banner
	}
	if len(u.Biography) > 0 {
		register["biography"] = u.Biography
	}
	if len(u.Location) > 0 {
		register["location"] = u.Location
	}
	if len(u.Name) > 0 {
		register["name"] = u.Name
	}
	if len(u.Surname) > 0 {
		register["surname"] = u.Surname
	}
	if len(u.Website) > 0 {
		register["website"] = u.Website
	}
	updtString := bson.M{
		"$set": register,
	}
	objID, _ := primitive.ObjectIDFromHex(ID)
	filter := bson.M{"_id": bson.M{"$eq": objID}} //mongo structure

	_, err := col.UpdateOne(ctx, filter, updtString)
	if err != nil {
		return false, err
	}
	return true, nil
}
