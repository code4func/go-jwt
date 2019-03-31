package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)


type User struct {
	ID          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"` // tag golang 
	Email       string             `json:"email" bson:"email"`
	Password    string             `json:"password" bson:"password"`
	DisplayName string             `json:"displayName" bson:"displayName"`
}
