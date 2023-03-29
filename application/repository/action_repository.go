package repository

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/hyperledger/fabric-gateway/pkg/client"
	"github.com/liuhdd/exam-cret/application/config"
	"github.com/liuhdd/exam-cret/application/models"
	"strconv"
)

type ActionRepository interface {
	AddAction(*models.ExamAction) error
	FindActionByActionID(string) (*models.ExamAction, error)
	FindActionsByStudentID(string) []*models.ExamAction
	FindActionsByExamAndStudentID(string, string) ([]*models.ExamAction, error)
	QueryAction(string) ([]*models.ExamAction, error)
}

type actionRepository struct {
	ActionRepository
	contract *client.Contract
}

var actionRepo *actionRepository

func NewActionRepository() ActionRepository {
	if actionRepo != nil {
		return actionRepo
	}
	contract := config.GetContract()
	actionRepo = &actionRepository{contract: contract}

	return actionRepo
}

func (a *actionRepository) AddAction(action *models.ExamAction) error {
	if action == nil {
		return errors.New("nil point of action")
	}
	_, err := a.contract.SubmitTransaction("UploadAction",
		action.ObjectType,
		action.ActionID,
		action.ExamID,
		action.StudentID,
		strconv.Itoa(int(action.ActionType)),
		strconv.FormatInt(action.ActionTime, 10),
		action.QuestionID,
		action.Answer,
	)
	if err != nil {
		return fmt.Errorf("failed to upload action: %s", err)
	}
	return nil

}

func (a *actionRepository) FindActionByActionID(id string) (*models.ExamAction, error) {
	if id == "" {
		return nil, errors.New("miss action id")
	}
	res, err := a.contract.SubmitTransaction("QueryActionByID", id)
	if err != nil {
		return nil, fmt.Errorf("failed to query action: %s", err)
	}
	action := &models.ExamAction{}
	err = json.Unmarshal(res, action)
	if err != nil {
		return nil, err
	}
	return action, nil
}

func (a *actionRepository) FindActionsByExamAndStudentID(examID, studentID string) ([]*models.ExamAction, error) {
	if examID == "" || studentID == "" {
		return nil, errors.New("miss examID and studentID")
	}
	bytes, err := a.contract.SubmitTransaction("QueryActionByExamAndStudentID", "exam_action", examID, studentID)
	if err != nil {
		return nil, fmt.Errorf("failed to query action: %s", err)
	}
	if bytes == nil {
		return nil, nil
	}
	var actions []*models.ExamAction
	err = json.Unmarshal(bytes, actions)
	if err != nil {
		return nil, err
	}
	return actions, nil
}

func (a *actionRepository) FindActionsByStudentID(s string) []*models.ExamAction {
	//TODO implement me
	panic("implement me")
}

func (a actionRepository) QueryAction(selector string) ([]*models.ExamAction, error) {
	bytes, err := a.contract.SubmitTransaction("QueryAction", selector)
	if err != nil {
		return nil, err
	}
	var actions []*models.ExamAction
	err = json.Unmarshal(bytes, &actions)
	if err != nil {
		return nil, err
	}
	return actions, nil
}
