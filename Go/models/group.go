package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Group struct {
	Id     primitive.ObjectID `json:"id,omitempty" bson:"id"`
	Name   string             `json:"name,omitempty" bson:"name"`
	UserId primitive.ObjectID `json:"userid,omitempty" bson:"userid"`
}
