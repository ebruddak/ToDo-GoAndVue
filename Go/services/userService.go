package services

import (
	"../models"
	"../repository"
	"../dtos"
	"../utils"
	"strconv"
	"errors"
	// "go.mongodb.org/mongo-driver/bson/primitive"
)

type DefaultUserService struct {
	Repo repository.UserRepository
}

type UserService interface {
	UserInsert(user models.User) (*dtos.TodoDTO, error)	
	Login(user dtos.LoginDTO) (string, error)	
}

func (t DefaultUserService) UserInsert(user models.User) (*dtos.TodoDTO, error) {
	var res dtos.TodoDTO
	if len(user.UserName) <= 3 {
		res.Status = false
		return &res, nil
	}
	user.SetPassword("1234")
	result, err := t.Repo.Insert(user)

	if err != nil || result == false {
		res.Status = false
		return &res, err
	}
	res = dtos.TodoDTO{Status: result}
	return &res, nil
}

func (t DefaultUserService) Login(user dtos.LoginDTO) (string , error) {
	result, err := t.Repo.Login(user)
    if err == nil {
		token, er := utils.GenerateJwt(strconv.Itoa(int(result.Id)))
		return token,er}
	return "", errors.New("failed") 				
}


func NewUserService(Repo repository.UserRepository) DefaultUserService {
	return DefaultUserService{Repo: Repo}
}