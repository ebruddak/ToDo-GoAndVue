package services

import (
	"errors"

	"../dtos"
	"../models"
	"../repository"
	"../utils"
	// "go.mongodb.org/mongo-driver/bson/primitive"
)

type DefaultUserService struct {
	Repo repository.UserRepository
}

type UserService interface {
	UserInsert(user dtos.RegisterDTO) (*dtos.TodoDTO, error)
	Login(user dtos.LoginDTO) (string, error)
	User(_id string) (*models.User, error)
}

func (t DefaultUserService) UserInsert(user dtos.RegisterDTO) (*dtos.TodoDTO, error) {
	var res dtos.TodoDTO
	var newUser models.User
	if len(user.UserName) <= 3 {
		res.Status = false
		return &res, nil
	}
	newUser.SetPassword(string(user.Password))
	newUser.Email = user.Email
	newUser.UserName = user.UserName
	result, err := t.Repo.Insert(newUser)

	if err != nil || result == false {
		res.Status = false
		return &res, err
	}
	res = dtos.TodoDTO{Status: result}
	return &res, nil
}

func (t DefaultUserService) Login(user dtos.LoginDTO) (string, error) {
	result, err := t.Repo.Login(user)

	if err != nil {
		return "", errors.New(err.Error())
	}
	token, er := utils.GenerateJwt(string(result.Id.String()))
	return token, er
}
func (t DefaultUserService) User(_id string) (*models.User, error) {
	result, err := t.Repo.User(_id)
	if err != nil {
		return &result, errors.New(err.Error())
	}
	return &result, nil
}

func NewUserService(Repo repository.UserRepository) DefaultUserService {
	return DefaultUserService{Repo: Repo}
}
