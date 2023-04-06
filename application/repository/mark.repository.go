package repository

import (
	"encoding/json"
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
}

type markRepository struct {
	db       *gorm.DB
	contract *client.Contract
}

func NewMarkRepository() MarkRepository {
	db := config.GetDB()
	db.AutoMigrate(&models.MarkAction{})
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

func (r *markRepository) FindMarkByQuestionIDFromDB(examID, studentID, questionID string) (*models.MarkAction, error) {
	var mark models.MarkAction
	tx := r.db.Where("exam_id=? and student_id=? and question_id=?", examID, studentID, questionID).
		Select("question_id, score, scorer, scored_time").
		First(&mark)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &mark, nil
}

func (r *markRepository) FindMarkByQuestionIDFromBC(examID, studentID, questionID string) (*models.MarkAction, error) {
	result, err := r.contract.EvaluateTransaction("GetQuestionScore", examID, studentID, questionID)
	if err != nil {
		return nil, err
	}
	var mark models.MarkAction
	err = json.Unmarshal(result, &mark)
	if err != nil {
		return nil, err
	}
	return &mark, nil
}
