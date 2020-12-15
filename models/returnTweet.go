package models

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//ReturnTweet is the structure that we will use to return tweets
type ReturnTweet struct {
	ID primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	UserID string `bson:"userid" json:"userID,omitempty"`
	Message string `bson:"message" json:"message,omitempty"`
	Date time.Time `bson:"date" json:"date,omitempty"`
}
