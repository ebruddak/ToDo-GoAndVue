package dtos

type LoginDTO struct {
	Email string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}