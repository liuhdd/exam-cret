package repository

import (
	"encoding/json"
	"errors"
	"strconv"

	"github.com/hyperledger/fabric-gateway/pkg/client"
	"github.com/liuhdd/exam-cret/application/config"
	"github.com/liuhdd/exam-cret/application/models"
	"gorm.io/gorm"
)

type MarkRepository interface {
	FindMarkByQuestionIDFromDB(string, string, string) (*models.MarkAction, error)
	FindMarkByQuestionIDFromBC(string, string, string) (*models.MarkAction, error)
	UploadMarkToDB(mark *models.MarkAction) error
	UploadMarkToBC(mark *models.MarkAction) error
	GetScores(examID, studentID string) ([]*models.MarkAction, error)
}

type markRepository struct {
	db       *gorm.DB
	contract *client.Contract
}

func NewMarkRepository() MarkRepository {
	db := config.GetDB()
	db.AutoMigrate(&models.MarkAction{})
	db.AutoMigrate(&models.Mark{})
	repo := &markRepository{db: db, contract: config.GetContract()}
	return repo
}

func (r *markRepository) UploadMarkToDB(mark *models.MarkAction) error {
	tx := r.db.Save(mark)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (r *markRepository) UploadMarkToBC(mark *models.MarkAction) error {

	_, err := r.contract.SubmitTransaction(
		"UploadMarkAction",
		mark.ActionID,
		mark.ExamID,
		mark.StudentID,
		mark.QuestionID,
		strconv.FormatUint(uint64(mark.Score), 10),
		strconv.FormatInt(mark.ScoredTime, 10),
		mark.Scorer)
	if err != nil {
		return err
	}

	return nil
}

func (r *markRepository) GetScores(examID, studentID string) ([]*models.MarkAction, error) {
	var marks []*models.MarkAction
	tx := r.db.Where("exam_id=? and student_id=?", examID, studentID).
		Group("question_id").
		Having("scored_time=Max(scored_time)").
		Select("question_id, score, scorer").
		Find(&marks)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return marks, tx.Error
}

func (r *markRepository) FindMarkByQuestionIDFromDB(examID, studentID, questionID string) (*models.MarkAction, error) {
	var mark models.MarkAction
	tx := r.db.Where("exam_id=? and student_id=? and question_id=?", examID, studentID, questionID).
		Select("question_id, score, scorer, scored_time").
		First(&mark)
	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, tx.Error
	}
	return &mark, nil
}

func (r *markRepository) FindMarkByQuestionIDFromBC(examID, studentID, questionID string) (*models.MarkAction, error) {
	result, err := r.contract.EvaluateTransaction("QuestionScore", examID, studentID, questionID)
	if err != nil {
		return nil, err
	}
	var marks []*models.MarkAction
	err = json.Unmarshal(result, &marks)
	if err != nil {
		return nil, err
	}
	if len(marks) == 0 {
		return nil, nil
	}
	return marks[0], nil
}
