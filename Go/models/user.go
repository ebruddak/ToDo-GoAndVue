package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id       primitive.ObjectID `json:"id,omitempty" bson:"_id"`
	UserName string             `json:"userName,omitempty" bson:"username"`
	Email    string             `json:"eMail,omitempty" bson:"email"`
	Password []byte             `json:"password,omitempty" bson:"password"`
}

func (user *User) SetPassword(password string) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	user.Password = hashedPassword
}
func (user *User) ComparePassword(password string) error {
	return bcrypt.CompareHashAndPassword(user.Password, []byte(password))
}
