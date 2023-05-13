package services

import (
	"context"
	"errors"
	"sync"

	"github.com/liuhdd/exam-cret/application/config"
	"github.com/liuhdd/exam-cret/application/models"
	"github.com/liuhdd/exam-cret/application/repository"
	"github.com/liuhdd/exam-cret/application/services/dto"
	"github.com/redis/go-redis/v9"
	log "github.com/sirupsen/logrus"
)

var as *actionService

type ActionService interface {
	QueryActionByID(string) (*models.ExamAction, error)
	UploadAction(action *models.ExamAction) error
	SelectActionByExamAndStudentID(string, string) ([]*models.ExamAction, error)
	QueryAction(string) ([]*models.ExamAction, error)
	ListActionInQuestion(string, string, string) ([]*models.ExamAction, error)
	UploadActions(actions *[]models.ExamAction) error
}

type actionService struct {
	AuthService
	actionRepo repository.ActionRepository
	client     *redis.Client
}

var once sync.Once

func NewActionService() ActionService {
	once.Do(func() {
		actionRepository := repository.NewActionRepository()
		c := config.GetRedisClient()
		as = &actionService{
			actionRepo: actionRepository,
			client:     c,
		}
	})
	return as
}

func (as *actionService) QueryActionByID(id string) (*models.ExamAction, error) {
	if id == "" {
		return nil, errors.New("miss param id")
	}
	action, err := as.actionRepo.FindActionByActionID(id)
	if err != nil {
		log.Printf("error in QueryActionByID: %s", err)
		return nil, err
	}
	return action, nil
}

func (as *actionService) UploadActions(actions *[]models.ExamAction) error {
	err := as.actionRepo.AddActions(actions)
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}
func (as *actionService) UploadAction(action *models.ExamAction) error {
	if action == nil {
		return errors.New("nil pointer to action")
	}
	ctx := context.Background()
	as.client.XAdd(ctx, &redis.XAddArgs{
		Stream: "action",
		ID:     "*",
		Values: map[string]interface{}{
			"object_type": action.ObjectType,
			"action_id":   action.ActionID,
			"exam_id":     action.ExamID,
			"student_id":  action.StudentID,
			"question_id": action.QuestionID,
			"answer":      action.Answer,
			"answer_time": action.ActionTime,
			"action_type": action.ActionType,
		},
	}).Result()

	return nil
}

func (as *actionService) SelectActionByExamAndStudentID(examID string, studentID string) ([]*models.ExamAction, error) {
	actions, err := as.actionRepo.FindActionsByExamAndStudentID(examID, studentID)
	if err != nil {
		log.Printf("error in find action: %s", err)
		return nil, err
	}
	return actions, err
}

func (as *actionService) QueryAction(query string) ([]*models.ExamAction, error) {
	bytes, err := as.actionRepo.QueryActionFromBC(query)
	if err != nil {
		log.Printf("error in QueryAction: %s", err)
		return nil, err
	}
	return bytes, nil
}

func (as *actionService) ListActionInQuestion(examID string, studentID string, questionID string) ([]*models.ExamAction, error) {
	actions, err := as.actionRepo.FindActionsByExamAndStudentID(examID, studentID)
	if err != nil {
		log.Printf("error in find action: %s", err)
		return nil, err
	}
	var result []*models.ExamAction
	for _, action := range actions {
		if action.QuestionID == questionID {
			result = append(result, action)
		}
	}
	return result, nil
}

func (as *actionService) ShowExamResult(examID string, studentID string) (*dto.ExamResult, error) {
	actions, err := as.actionRepo.GetAnswersFromDB(examID, studentID)
	if err != nil {
		log.Printf("error in find action: %s", err)
		return nil, err
	}
	var result []*dto.QuestionResult
	for _, action := range actions {
		result = append(result, &dto.QuestionResult{
			QuestionID: action.QuestionID,
			Answer:     action.Answer,
		})
	}
	return &dto.ExamResult{
		ExamID:    examID,
		StudentID: studentID,
		Questions: result,
	}, nil
}
