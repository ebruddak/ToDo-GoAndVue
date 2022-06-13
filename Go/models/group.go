package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Group struct {
	Id      primitive.ObjectID 				`json:"id,omitempty"`
	Name    string             				`json:"name,omitempty"`
	UserId  primitive.ObjectID              `json:"userId,omitempty"`
}