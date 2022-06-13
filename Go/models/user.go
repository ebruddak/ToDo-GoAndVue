package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id          primitive.ObjectID 				`json:"id,omitempty"`
	UserName    string             				`json:"userName,omitempty"`
	EMail       string             				`json:"eMail,omitempty"`
	Password    string             				`json:"password,omitempty"`
}