package services

import (
	"errors"
	log "github.com/sirupsen/logrus"

	"github.com/liuhdd/exam-cret/application/config"
	"github.com/liuhdd/exam-cret/application/models"
	"github.com/liuhdd/exam-cret/application/repository"
	"github.com/liuhdd/exam-cret/application/services/dto"
	"gorm.io/gorm"
)

type MarkService interface {
	FindMarkByQuestionID(examID, studentID, questionID string) (*dto.Score, error)
	VerificationQuestionScore(*dto.Question) (*dto.Score, bool, error)
	UploadMarkAction(mark *models.MarkAction) error
	GetScores(examID, studentID string) ([]*dto.Score, error)
}

type markService struct {
	MarkService
	db       *gorm.DB
	markRepo repository.MarkRepository
}

func NewMarkService() MarkService {
	markRepo := repository.NewMarkRepository()
	db := config.GetDB()
	return &markService{markRepo: markRepo, db: db}
}

func (s *markService) UploadMarkAction(mark *models.MarkAction) error {
	if mark == nil {
		log.Printf("nil pointer to mark")
		return errors.New("nil pointer to mark")
	}
	err := s.markRepo.UploadMarkToDB(mark)
	if err != nil {
		log.Printf("failed to upload mark to database: %s", err)
		return err
	}

	err = s.markRepo.UploadMarkToBC(mark)
	if err != nil {
		log.Printf("failed to upload mark to blockchain: %s", err)
		return err
	}
	m := models.Mark{
		ExamID:     mark.ExamID,
		StudentID:  mark.StudentID,
		QuestionID: mark.QuestionID,
	}
	s.db.Table("marks").Where("exam_id = ? and student_id = ? and question_id = ?",
		m.ExamID, m.StudentID, m.QuestionID).Select("answer").
		Scan(&m)
	m.Score = mark.Score
	s.db.Save(&m)
	var grade int
	s.db.Table("marks").
		Select("SUM(score) as grade").Where("exam_id = ? and student_id = ?", mark.ExamID, mark.StudentID).
		Scan(&grade)
	s.db.Save(&models.ExamRecord{
		ExamID:    mark.ExamID,
		StudentID: mark.StudentID,
		Grade:     grade,
	})
	return nil
}

func (s *markService) GetScores(examID, studentID string) ([]*dto.Score, error) {
	marks, err := s.markRepo.GetScores(examID, studentID)
	if err != nil {
		log.Printf("failed to get socres: %s", err)
		return nil, err
	}
	var scores []*dto.Score
	for _, mark := range marks {
		scores = append(scores, &dto.Score{
			QuestionID: mark.QuestionID,
			Score:      mark.Score,
			ScoredBy:   mark.Scorer,
			ScoredTime: mark.ScoredTime,
		})
	}
	return scores, nil
}
func (s *markService) FindMarkByQuestionID(examID, studentID, questionID string) (*dto.Score, error) {
	mark, err := s.markRepo.FindMarkByQuestionIDFromDB(examID, studentID, questionID)
	if err != nil {
		return nil, err
	}
	if mark == nil {
		return nil, nil
	}
	return &dto.Score{
		QuestionID: mark.QuestionID,
		Score:      mark.Score,
		ScoredBy:   mark.Scorer,
		ScoredTime: mark.ScoredTime,
	}, nil
}

func (s *markService) VerificationQuestionScore(question *dto.Question) (*dto.Score, bool, error) {
	mark, err := s.markRepo.FindMarkByQuestionIDFromBC(question.ExamID, question.StudentID, question.QuestionID)
	if err != nil {
		return nil, false, err
	}
	if mark == nil {
		return nil, false, nil
	}
	socre := &dto.Score{
		QuestionID: mark.QuestionID,
		Score:      mark.Score,
		ScoredBy:   mark.Scorer,
		ScoredTime: mark.ScoredTime,
	}
	if mark.Score == question.Score {
		return socre, true, nil
	}
	return socre, false, nil
}
