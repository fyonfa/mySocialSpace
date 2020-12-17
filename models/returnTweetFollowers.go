package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)
//ReturnTweetFollowers is the structure with the one we give back the tweets
type ReturnTweetFollowers struct {
	ID             primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	UserRelationID string             `bson:"userrelationid" json:"userRelationID,omitempty"`
	UserID         string             `bson:"userid" json:"userID,omitempty"`
	Tweet          struct {
		Message string    `bson:"message" json:"message,omitempty"`
		Date    time.Time `bson:"date" json:"date,omitempty"`
		ID      string    `bson:"_id" json:"_id,omitempty"`
	}
}
