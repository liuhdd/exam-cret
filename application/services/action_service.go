package services

import (
	"errors"
	"github.com/liuhdd/exam-cret/application/models"
	"github.com/liuhdd/exam-cret/application/repository"
	"log"
)

var as *actionService

type ActionService interface {
	QueryActionByID(string) (*models.ExamAction, error)
	UploadAction(action *models.ExamAction) error
}

type actionService struct {
	AuthService
	actionRepo repository.ActionRepository
}

func NewActionService() ActionService {
	if as != nil {
		return as
	}

	actionRepository := repository.NewActionRepository()
	as = &actionService{actionRepo: actionRepository}
	return as
}

func (a *actionService) QueryActionByID(id string) (*models.ExamAction, error) {
	if id == "" {
		return nil, errors.New("miss param id")
	}
	action, err := a.actionRepo.FindActionByActionID(id)
	if err != nil {
		log.Fatalf("error in QueryActionByID: %s", err)
		return nil, err
	}
	return action, nil
}

func (a *actionService) UploadAction(action *models.ExamAction) error {
	if action == nil {
		return errors.New("nil pointer to action")
	}
	err := a.actionRepo.AddAction(action)
	if err != nil {
		log.Fatalf("failed to upload action: %s", err)
		return err
	}
	return nil
}
