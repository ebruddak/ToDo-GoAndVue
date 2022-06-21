package repository

import (
	"context"
	"log"
	"time"

	"github.com/ebruddak/ToDo-GoAndVue/tree/main/Go/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type GroupRepositoryDB struct {
	GroupCollection *mongo.Collection
}

type GroupRepository interface {
	Insert(group models.Group) (bool, error)
	Update(group models.Group) (bool, error)
	GetAll(id primitive.ObjectID) ([]models.Group, error)
	Get(id primitive.ObjectID) (models.Group, error)
	Delete(id primitive.ObjectID) (bool, error)
}

func (t GroupRepositoryDB) Insert(group models.Group) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	group.Id = primitive.NewObjectID()
	result, err := t.GroupCollection.InsertOne(ctx, group)

	if result.InsertedID == nil || err != nil {
		return false, err
	}
	return true, nil
}
func (t GroupRepositoryDB) Update(group models.Group) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	id := primitive.ObjectID(group.Id)
	filter := bson.M{"id": bson.M{"$eq": id}}
	update := bson.M{"$set": bson.M{"name": group.Name, "userid": group.UserId}}
	result, err := t.GroupCollection.UpdateOne(ctx, filter, update)
	if err != nil || result == nil {
		return false, err
	}

	return true, nil
}
func (t GroupRepositoryDB) Get(id primitive.ObjectID) (models.Group, error) {
	var group models.Group

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result, err := t.GroupCollection.Find(ctx, bson.M{"id": id})

	if err != nil {
		log.Fatalln(err)
		return group, err
	}

	for result.Next(ctx) {
		if err := result.Decode(&group); err != nil {
			log.Fatalln(err)
			return group, err
		}
	}
	return group, nil
}
func (t GroupRepositoryDB) GetAll(id primitive.ObjectID) ([]models.Group, error) {
	var group models.Group
	var groups []models.Group

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := t.GroupCollection.Find(ctx, bson.M{"userid": id})
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	for result.Next(ctx) {
		if err := result.Decode(&group); err != nil {
			log.Fatalln(err)

			return nil, err
		}

		groups = append(groups, group)
	}

	return groups, nil
}

func (t GroupRepositoryDB) Delete(id primitive.ObjectID) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := t.GroupCollection.DeleteOne(ctx, bson.M{"id": id})

	if err != nil || result.DeletedCount <= 0 {
		return false, err
	}
	return true, nil
}

func NewGroupRepositoryDb(dbClient *mongo.Collection) GroupRepositoryDB {
	return GroupRepositoryDB{GroupCollection: dbClient}
}
