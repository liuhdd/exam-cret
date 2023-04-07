package repository

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric-gateway/pkg/client"
	"github.com/liuhdd/exam-cret/application/config"
	"github.com/liuhdd/exam-cret/application/models"
	"gorm.io/gorm"
)

type ActionRepository interface {
	AddAction(*models.ExamAction) error
	FindActionByActionID(string) (*models.ExamAction, error)
	FindActionsByStudentID(string) []*models.ExamAction
	FindActionsByExamAndStudentID(string, string) ([]*models.ExamAction, error)
	QueryActionFromBC(string) ([]*models.ExamAction, error)
	GetQuestionAnswerFromDB(string, string, string) (*models.ExamAction, error)
	GetAnswersFromDB(string, string) ([]*models.ExamAction, error)
}

type actionRepository struct {
	ActionRepository
	contract *client.Contract
	db       *gorm.DB
}

var actionRepo *actionRepository

func NewActionRepository() ActionRepository {
	if actionRepo != nil {
		return actionRepo
	}
	contract := config.GetContract()
	db := config.GetDB()
	db.AutoMigrate(&models.ExamAction{})
	actionRepo = &actionRepository{contract: contract, db: db}
	return actionRepo
}

func (a *actionRepository) AddAction(action *models.ExamAction) error {
	if action == nil {
		return errors.New("nil point of action")
	}
	tx := a.db.Save(action)

	if tx.Error != nil {
		return tx.Error
	}
	_, err := a.contract.SubmitTransaction("UploadAction",
		action.ActionID,
		action.ExamID,
		action.StudentID,
		strconv.Itoa(int(action.ActionType)),
		strconv.FormatInt(action.ActionTime, 10),
		action.QuestionID,
		action.Answer,
	)

	if err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to upload action: %s", err)
	}
	tx.Commit()
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
	fmt.Println(bytes)
	var actions []*models.ExamAction
	err = json.Unmarshal(bytes, &actions)
	if err != nil {
		return nil, err
	}
	return actions, nil
}

func (a *actionRepository) FindActionsByStudentID(s string) []*models.ExamAction {
	//TODO implement me
	panic("implement me")
}

func (a *actionRepository) QueryActionFromBC(selector string) ([]*models.ExamAction, error) {
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

func (a *actionRepository) GetAnswersFromDB(examID, studentID string) ([]*models.ExamAction, error) {
	tx := a.db.Where("exam_id = ? AND student_id = ? ", examID, studentID).
		Group("question_id").
		Having("action_time = MAX(action_time)").
		Select("question_id, answer").
		Find(&models.ExamAction{})
	if tx.Error != nil {
		return nil, tx.Error
	}
	var actions []*models.ExamAction
	tx.Scan(&actions)
	return actions, nil
}

func (a *actionRepository) GetQuestionAnswerFromDB(examID, studentID, questionID string) (*models.ExamAction, error) {
	tx := a.db.Where("exam_id = ? AND student_id = ? AND question_id = ?", examID, studentID, questionID).
		Order("action_time desc").
		Limit(1).
		Find(&models.ExamAction{})
	if tx.Error != nil {
		return nil, tx.Error
	}
	var action models.ExamAction
	tx.Scan(&action)
	return &action, nil
}
