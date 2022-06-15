package models

import (
	"golang.org/x/crypto/bcrypt"
	// "go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id          uint   							`json:"id,omitempty"`
	UserName    string             				`json:"userName,omitempty"`
	EMail       string             				`json:"eMail,omitempty"`
	Password    []byte             				`json:"password,omitempty"`
}
func (user *User) SetPassword(password string) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	user.Password = hashedPassword
}
func (user *User) ComparePassword(password string) error {
	return bcrypt.CompareHashAndPassword(user.Password, []byte(password))
}
