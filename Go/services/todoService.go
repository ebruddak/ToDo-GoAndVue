package services

import (
	"../dtos"
	"../models"
	"../repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DefaultTodoService struct {
	Repo repository.TodoRepository
}

type TodoService interface {
	TodoInsert(todo models.Todo) (*dtos.TodoDTO, error)
	TodoGetAll() ([]models.Todo, error)
	TodoGetAllCompated() ([]models.Todo, error)
	TodoDelete(id primitive.ObjectID) (bool, error)
	GetTodo(id primitive.ObjectID) (models.Todo, error)
	Complete(id primitive.ObjectID) (bool, error)
	TodoUpdate(todo models.Todo) (bool, error)
}

func (t DefaultTodoService) TodoUpdate(todo models.Todo) (bool, error) {

	result, err := t.Repo.Update(todo)

	if err != nil || result == false {
		return result, err
	}
	return result, nil
}
func (t DefaultTodoService) Complete(id primitive.ObjectID) (bool, error) {

	result, err := t.Repo.Complete(id)

	if err != nil || result == false {
		return result, err
	}
	return result, nil
}
func (t DefaultTodoService) TodoInsert(todo models.Todo) (*dtos.TodoDTO, error) {
	var res dtos.TodoDTO
	if len(todo.Title) <= 3 {
		res.Status = false
		return &res, nil
	}

	result, err := t.Repo.Insert(todo)

	if err != nil || result == false {
		res.Status = false
		return &res, err
	}

	res = dtos.TodoDTO{Status: result}
	return &res, nil
}

func (t DefaultTodoService) TodoGetAll() ([]models.Todo, error) {
	result, err := t.Repo.GetAll()
	if err != nil {
		return nil, err
	}
	return result, nil
}
func (t DefaultTodoService) TodoGetAllCompated() ([]models.Todo, error) {
	result, err := t.Repo.GetAllComplated()
	if err != nil {
		return nil, err
	}
	return result, nil
}
func (t DefaultTodoService) GetTodo(id primitive.ObjectID) (models.Todo, error) {
	result, err := t.Repo.Get(id)
	if err != nil {
		return result, err
	}
	return result, nil
}
func (t DefaultTodoService) TodoDelete(id primitive.ObjectID) (bool, error) {
	result, err := t.Repo.Delete(id)

	if err != nil || result == false {
		return false, err
	}

	return true, nil
}

func NewTodoService(Repo repository.TodoRepository) DefaultTodoService {
	return DefaultTodoService{Repo: Repo}
}
