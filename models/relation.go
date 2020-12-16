package models

//Relation is a model to store the relation from one user to other
type Relation struct {
	UserID         string `bson:"userid" json:"userID"` //when click in follow button it will record the first data my userID and second the userI follow
	UserRelationID string `bson:"userrelationid" json:"userRelationID"`
}
