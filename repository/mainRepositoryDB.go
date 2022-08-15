package repository

import (
	"context"
	"forLearnCurrent/models"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"time"
)

//go:generate mockgen -destination=../mocks/repository/mockMainService.go -package=services forLearnCurrent/repository MainRepository
type MainRepositoryDB struct {
	MainCollection *mongo.Collection
}

type MainRepository interface {
	Insert(main models.Main) (bool, error)
	GetAll() ([]models.Main, error)
	GetById(id primitive.ObjectID) (models.Main, error)
	Delete(id primitive.ObjectID) (bool, error)
}

func (t MainRepositoryDB) Insert(main models.Main) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	main.Id = primitive.NewObjectID()
	result, err := t.MainCollection.InsertOne(ctx, main)

	if result.InsertedID == nil || err != nil {
		errors.New("Error occured when adding new main")
		return false, err
	}

	return true, nil
}

func (t MainRepositoryDB) GetAll() ([]models.Main, error) {
	var main models.Main
	var mains []models.Main

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := t.MainCollection.Find(ctx, bson.M{})

	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	for result.Next(ctx) {
		if err := result.Decode(&main); err != nil {
			log.Fatalln(err)
			return nil, err
		}

		mains = append(mains, main)
	}

	return mains, nil
}

func (t MainRepositoryDB) GetById(id primitive.ObjectID) (models.Main, error) {
	var main models.Main

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := t.MainCollection.FindOne(ctx, bson.M{"id": id}).Decode(&main)

	if err != nil {
		log.Fatalln(err)
		return main, err
	}

	return main, nil
}

func (t MainRepositoryDB) Delete(id primitive.ObjectID) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := t.MainCollection.DeleteOne(ctx, bson.M{"id": id})

	if err != nil || result.DeletedCount <= 0 {
		log.Fatalln(err)
		return false, err
	}

	return true, nil
}

func NewMainRepositoryDB(dbClient *mongo.Collection) MainRepositoryDB {
	return MainRepositoryDB{MainCollection: dbClient}
}
