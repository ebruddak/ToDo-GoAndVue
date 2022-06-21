package repository

import (
	"context"
	"errors"
	"time"

	"github.com/ebruddak/ToDo-GoAndVue/tree/main/Go/dtos"
	"github.com/ebruddak/ToDo-GoAndVue/tree/main/Go/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	var user models.User

	cnv, _ := primitive.ObjectIDFromHex(_id)

	result, err := t.UserCollection.Find(ctx, bson.M{"id": cnv})

	if err != nil {
		return user, err
	}

	for result.Next(ctx) {
		if err := result.Decode(&user); err != nil {

			return user, nil
		}
	}

	return user, nil

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
	userResult := models.User{}

	err := t.UserCollection.FindOne(ctx, bson.D{{Key: "email", Value: user.Email}}).Decode(&userResult)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return userResult, errors.New("User not found")
		}
	}

	erp := userResult.ComparePassword(user.Password)

	if erp != nil {
		return userResult, errors.New("Incorrect Pasword")
	}

	return userResult, nil

}
func NewUserRepositoryDb(dbClient *mongo.Collection) UserRepositoryDB {
	return UserRepositoryDB{UserCollection: dbClient}
}
