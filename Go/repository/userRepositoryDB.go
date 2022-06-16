package repository

import (
	"context"
	"errors"

	"../dtos"
	"../models"

	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepositoryDB struct {
	UserCollection *mongo.Collection
}

type UserRepository interface {
	Insert(user models.User) (bool, error)
	Login(user dtos.LoginDTO) (models.User, error)
	User(_id string) (models.User, error)
}

func (t UserRepositoryDB) User(_id string) (models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	filter := bson.M{"id": _id}
	res := models.User{}
	errr := t.UserCollection.FindOne(ctx, filter).Decode(&res)
	if errr != mongo.ErrNoDocuments {
		return res, errors.New(errr.Error())
	}

	return res, nil
}
func (t UserRepositoryDB) Insert(user models.User) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	filter := bson.M{"email": user.Email}
	res := models.User{}
	errr := t.UserCollection.FindOne(ctx, filter).Decode(&res)
	if errr != mongo.ErrNoDocuments {
		return false, errors.New(errr.Error())
	}
	user.Id = primitive.NewObjectID()
	result, err := t.UserCollection.InsertOne(ctx, user)

	if result.InsertedID == nil || err != nil {
		// errors.New("failed add")
		return false, err
	}
	return true, nil
}

func (t UserRepositoryDB) Login(user dtos.LoginDTO) (models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"email": user.Email}
	result := models.User{}
	err := t.UserCollection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			// return models.User{}, nil
		}

		return models.User{}, errors.New("User Not Found")
	}
	if err := result.ComparePassword(user.Password); err != nil {
		return models.User{}, errors.New("Incorrect Pasword")
	}

	return result, nil

}
func NewUserRepositoryDb(dbClient *mongo.Collection) UserRepositoryDB {
	return UserRepositoryDB{UserCollection: dbClient}
}
