package models

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Todo struct {
	Id       primitive.ObjectID 				`json:"id,omitempty"`
	Title    string             				`json:"title,omitempty"`
	Content  string             				`json:"content,omitempty"`
	Priority string             				`json:"priority,omitempty"`
	UserId   primitive.ObjectID                 `json:"userId,omitempty"`
	GroupId  primitive.ObjectID                 `json:"groupId,omitempty"`
	DueDate  time.Time          				`json:"dueDate,omitempty"`
}