package dtos

type LoginDTO struct {
	Email    string `json:"email,omitempty" bson:"email"`
	Password string `json:"password,omitempty" bson:"password"`
}
