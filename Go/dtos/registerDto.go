package dtos

type RegisterDTO struct {
	UserName string `json:"userName,omitempty" bson:"username"`
	Email    string `json:"eMail,omitempty" bson:"email"`
	Password string `json:"password,omitempty" bson:"password"`
}
