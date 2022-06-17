package repository

import (
	"context"

	"../models"

	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TodoRepositoryDB struct {
	TodoCollection *mongo.Collection
}

type TodoRepository interface {
	Insert(todo models.Todo) (bool, error)
	GetAll(id primitive.ObjectID) ([]models.Todo, error)
	GetAllComplated(id primitive.ObjectID) ([]models.Todo, error)
	Delete(id primitive.ObjectID) (bool, error)
	Get(id primitive.ObjectID) (models.Todo, error)
	Update(todo models.Todo) (bool, error)
	Complete(id primitive.ObjectID) (bool, error)
}

func (t TodoRepositoryDB) Insert(todo models.Todo) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	todo.Id = primitive.NewObjectID()
	result, err := t.TodoCollection.InsertOne(ctx, todo)

	if result.InsertedID == nil || err != nil {
		return false, err
	}
	return true, nil
}

func (t TodoRepositoryDB) Update(todo models.Todo) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	id := primitive.ObjectID(todo.Id)
	filter := bson.M{"id": bson.M{"$eq": id}}
	update := bson.M{"$set": bson.M{"title": todo.Title, "content": todo.Content, "priority": todo.Priority, "groupId": todo.GroupId, "userId": todo.UserId}}
	result, err := t.TodoCollection.UpdateOne(ctx, filter, update)
	if err != nil || result == nil {
		return false, err
	}

	return true, nil
}
func (t TodoRepositoryDB) Complete(id primitive.ObjectID) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	filter := bson.M{"id": bson.M{"$eq": id}}
	update := bson.M{"$set": bson.M{"state": false}}
	result, err := t.TodoCollection.UpdateOne(ctx, filter, update)
	if err != nil || result == nil {
		return false, err
	}
	return true, nil
}
func (t TodoRepositoryDB) Get(id primitive.ObjectID) (models.Todo, error) {
	var todo models.Todo

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result, err := t.TodoCollection.Find(ctx, bson.M{"id": id})

	if err != nil {
		log.Fatalln(err)
		return todo, err
	}

	for result.Next(ctx) {
		if err := result.Decode(&todo); err != nil {
			log.Fatalln(err)
			return todo, err
		}
	}
	return todo, nil
}
func (t TodoRepositoryDB) GetAll(id primitive.ObjectID) ([]models.Todo, error) {
	var todo models.Todo
	var todos []models.Todo

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := t.TodoCollection.Find(ctx, bson.M{"userid": id, "state": true})

	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	for result.Next(ctx) {
		if err := result.Decode(&todo); err != nil {
			log.Fatalln(err)
			return nil, err
		}
		todos = append(todos, todo)
	}
	return todos, nil
}
func (t TodoRepositoryDB) GetAllComplated(id primitive.ObjectID) ([]models.Todo, error) {
	var todo models.Todo
	var todos []models.Todo

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := t.TodoCollection.Find(ctx, bson.M{"userid": id, "state": false})

	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	for result.Next(ctx) {
		if err := result.Decode(&todo); err != nil {
			log.Fatalln(err)
			return nil, err
		}
		todos = append(todos, todo)
	}
	return todos, nil
}
func (t TodoRepositoryDB) Delete(id primitive.ObjectID) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := t.TodoCollection.DeleteOne(ctx, bson.M{"id": id})

	if err != nil || result.DeletedCount <= 0 {
		return false, err
	}
	return true, nil
}

func NewTodoRepositoryDb(dbClient *mongo.Collection) TodoRepositoryDB {
	return TodoRepositoryDB{TodoCollection: dbClient}
}
