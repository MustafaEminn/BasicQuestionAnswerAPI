package services

import (
	"forLearnCurrent/dto"
	"forLearnCurrent/models"
	"forLearnCurrent/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//go:generate mockgen -destination=../mocks/service/mockMainService.go -package=services forLearnCurrent/services MainService
type DefaultMainService struct {
	Repo repository.MainRepository
}

type MainService interface {
	MainInsert(main models.Main) (*dto.MainDTO, error)
	MainGetById(id primitive.ObjectID) (models.Main, error)
	MainGetAll() ([]models.Main, error)
	MainDelete(id primitive.ObjectID) (bool, error)
}

func (t DefaultMainService) MainInsert(main models.Main) (*dto.MainDTO, error) {
	var res dto.MainDTO

	if len(main.Title) <= 2 || len(main.Content) < 10 {
		res.Status = false
		return &res, nil
	}

	main.Comments = []models.Comment{}
	main.Upvote = 1
	main.Downvote = 0

	result, err := t.Repo.Insert(main)

	if err != nil || result == false {
		res.Status = false
		return &res, err
	}

	res = dto.MainDTO{Status: result}

	return &res, nil
}

func (t DefaultMainService) MainGetById(id primitive.ObjectID) (models.Main, error) {
	var main = models.Main{}

	result, err := t.Repo.GetById(id)

	if err != nil {
		return main, err
	}

	return result, nil
}

func (t DefaultMainService) MainGetAll() ([]models.Main, error) {
	result, err := t.Repo.GetAll()

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (t DefaultMainService) MainDelete(id primitive.ObjectID) (bool, error) {
	result, err := t.Repo.Delete(id)

	if err != nil || !result {
		return false, err
	}

	return true, nil
}

func NewMainService(Repo repository.MainRepository) DefaultMainService {
	return DefaultMainService{Repo: Repo}
}
