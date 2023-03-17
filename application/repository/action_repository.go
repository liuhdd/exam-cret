package repository

import (
	"github.com/liuhdd/exam-cret/application/models"
)

type ActionRepository interface {
	AddAction(*models.ExamAction) error
	FindActionByActionID(string) *models.ExamAction
	FindActionsByStudentID(string) []*models.ExamAction
	FindActionsByExamID(string) []*models.ExamAction
	QueryAction(string) []*models.ExamAction
}

type actionRepository struct {
}

func NewActionRepository() ActionRepository {
	return nil

}
