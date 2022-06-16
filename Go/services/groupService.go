package services

import (
	"../dtos"
	"../models"
	"../repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DefaultGroupService struct {
	Repo repository.GroupRepository
}

type GroupService interface {
	GroupInsert(group models.Group) (*dtos.GroupDTO, error)
	GroupUpdate(group models.Group) (bool, error)
	GroupGetAll() ([]models.Group, error)
	GroupDelete(id primitive.ObjectID) (bool, error)
	GetGroup(id primitive.ObjectID) (models.Group, error)
}

func (t DefaultGroupService) GroupInsert(group models.Group) (*dtos.GroupDTO, error) {
	var res dtos.GroupDTO

	result, err := t.Repo.Insert(group)

	if err != nil || result == false {
		res.Status = false
		return &res, err
	}

	res = dtos.GroupDTO{Status: result}
	return &res, nil
}
func (t DefaultGroupService) GroupUpdate(group models.Group) (bool, error) {

	result, err := t.Repo.Update(group)

	if err != nil || result == false {
		return result, err
	}
	return result, nil
}
func (t DefaultGroupService) GroupGetAll() ([]models.Group, error) {
	result, err := t.Repo.GetAll()
	if err != nil {
		return nil, err
	}
	return result, nil
}
func (t DefaultGroupService) GetGroup(id primitive.ObjectID) (models.Group, error) {
	result, err := t.Repo.Get(id)
	if err != nil {
		return result, err
	}
	return result, nil
}
func (t DefaultGroupService) GroupDelete(id primitive.ObjectID) (bool, error) {
	result, err := t.Repo.Delete(id)

	if err != nil || result == false {
		return false, err
	}

	return true, nil
}

func NewGroupService(Repo repository.GroupRepository) DefaultGroupService {
	return DefaultGroupService{Repo: Repo}
}
