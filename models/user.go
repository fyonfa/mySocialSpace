package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

//User is the model of the user for the MongoDB. profile, this is for the user register
type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`          //id from mongo, it is a binary object, slice of bits, all this is standard from mongo
	Name      string             `bson:"name" json:"name,omitempty"`       //on the database the name will be "name" and json the same
	Surname   string             `bson:"surname" json:"surname,omitempty"` //dont give it back in json if you dont find it, omitempty, ones are the data out and others data in
	BirthDate time.Time          `bson:"birthDate" json:"birthDate,omitempty"`
	Email     string             `bson:"email" json:"email"`                 //not emit empty, always returns email
	Password  string             `bson:"password" json:"password,omitempty"` //***password allays is omitempty, never will give back a password trough the browser, never give back an endpoint trough the web
	Avatar    string             `bson:"avatar" json:"avatar,omitempty"`
	Banner    string             `bson:"banner" json:"banner,omitempty"`
	Biography string             `bson:"biography" json:"biography,omitempty"`
	Location  string             `bson:"location" json:"location,omitempty"`
	Website   string             `bson:"website" json:"website,omitempty"`
}
