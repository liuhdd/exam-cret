package services

import (
	"github.com/liuhdd/exam-cret/application/repository"
	"github.com/liuhdd/exam-cret/application/services/dto"
)

type MarkService interface {
	FindMarkByQuestionID (examID, studentID, questionID string) (*dto.Score, error)
	VerificationQuestionScore (*dto.Question) (*dto.Score, bool,  error)
}

type markService struct {
	MarkService
	markRepo repository.MarkRepository
}

func NewMarkService() MarkService {
	markRepo := repository.NewMarkRepository()
	return &markService{markRepo: markRepo}
}

func (s *markService) FindMarkByQuestionID(examID, studentID, questionID string) (*dto.Score, error) {
	mark, err := s.markRepo.FindMarkByQestionIDFromDB(examID, studentID, questionID)
	if err != nil {
		return nil, err
	}
	return &dto.Score{
		QuestionID: mark.QuestionID,
		Score: mark.Score,
		}, nil
}

func (a *markService) VerificationQuestionScore(question *dto.Question) (*dto.Score, bool, error) {
	mark, err := a.markRepo.FindMarkByQestionIDFromBC(question.ExamID, question.StudentID, question.QuestionID)
	if err != nil {
		return nil, false, err
	}
	socre := &dto.Score{
		QuestionID: mark.QuestionID,
		Score: mark.Score,
		ScoredBy: mark.Scorer,
		ScoredTime: mark.ScoredTime,
	}
	if mark.Score == question.Score {
		return socre, true, nil
	}
	return socre, false, nil
}