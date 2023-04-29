package services

import (
	"github.com/liuhdd/exam-cret/application/models"
	"github.com/liuhdd/exam-cret/application/repository"
	"github.com/liuhdd/exam-cret/application/services/dto"
	log "github.com/sirupsen/logrus"
)

type ExamService interface {
	FindExamResultByExamIDAndStudentID(examID, studentID string) (*dto.ExamResult, error)
	VerifyExamResults(*dto.ExamResult) (*dto.ExamProcess, bool, error)
	FindExamByID(id string) (*models.Exam, error)
	FindExamsByStudentID(id string) ([]*models.Exam, error)
	FindExamsByName(name string) ([]*models.Exam, error)
	FindExamsByTime(beginTime int64, endTime int64) ([]*models.Exam, error)
	SaveExam(exam *models.Exam) error
	FindExamRecordByID(id string) (*models.ExamRecord, error)
	FindExamRecordsByStudentID(id string) ([]*models.ExamRecord, error)
	FindExamRecordsByExamID(id string) ([]*models.ExamRecord, error)
	SaveExamRecord(examRecord *models.ExamRecord) error
	FindExamRecordsByExamIDAndStudentID(examID string, studentID string) ([]*models.ExamRecord, error)
}

type examService struct {
	examRepo repository.ExamRepository
	actionService ActionService
	markService   MarkService
}

func NewExamService() ExamService {
	return &examService{
		actionService: NewActionService(),
		markService:   NewMarkService(),
	}
}

func (s *examService) FindExamByID(id string) (*models.Exam, error) {
	return s.examRepo.GetExamByID(id)
}

func (s *examService) FindExamsByName(name string) ([]*models.Exam, error) {
	return s.examRepo.GetExamsByName(name)
}

func (s *examService) FindExamsByTime(beginTime int64, endTime int64) ([]*models.Exam, error) {
	return s.examRepo.GetExamsByTime(beginTime, endTime)
}

func (s *examService) SaveExam(exam *models.Exam) error {
	return s.examRepo.SaveExam(exam)
}

func (s *examService) FindExamRecordByID(id string) (*models.ExamRecord, error) {
	return s.examRepo.GetExamRecordByID(id)
}

func (s *examService) FindExamRecordsByStudentID(id string) ([]*models.ExamRecord, error) {
	return s.examRepo.GetExamRecordsByStudentID(id)
}

func (s *examService) FindExamRecordsByExamID(id string) ([]*models.ExamRecord, error) {
	return s.examRepo.GetExamRecordsByExamID(id)
}

func (s *examService) SaveExamRecord(examRecord *models.ExamRecord) error {
	return s.examRepo.SaveExamRecord(examRecord)
}

func (s *examService) FindExamRecordsByExamIDAndStudentID(examID string, studentID string) ([]*models.ExamRecord, error) {
	return s.examRepo.GetExamRecordsByExamIDAndStudentID(examID, studentID)
}

func (s *examService) FindExamsByStudentID(id string) ([]*models.Exam, error) {
	return s.examRepo.GetExamsByStudentID(id)
}
func (s *examService) FindExamResultByExamIDAndStudentID(examID, studentID string) (*dto.ExamResult, error) {
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
	scores, err := s.markService.GetScores(examID, studentID)
	if err != nil {
		log.Printf("%s", err)
		return nil, err
	}
	for _, score := range scores {
		for _, question := range result {
			if score.QuestionID == question.QuestionID {
				question.Score = score.Score
				break
			}
		}
	}
	return &dto.ExamResult{
		ExamID:    examID,
		StudentID: studentID,
		Questions: result,
	}, nil
}

func (s *examService) VerifyExamResults(result *dto.ExamResult) (*dto.ExamProcess, bool, error) {
	examID := result.ExamID
	studentID := result.StudentID

	actions, err := s.actionService.SelectActionByExamAndStudentID(examID, studentID)
	if err != nil {
		log.Error(err)
		return nil, false, err
	}

	process := &dto.ExamProcess{
		ExamID:    examID,
		StudentID: studentID,
	}

	var correct bool = true

	for _, res := range result.Questions {

		qusInfo := &dto.QuestionInfo{QuestionID: res.QuestionID}

		for _, action := range actions {

			if res.QuestionID == action.QuestionID {
				actInfo := &dto.ActionInfo{
					ActionID:   action.ActionID,
					Answer:     action.Answer,
					ActionTime: action.ActionTime,
				}
				qusInfo.Actions = append(qusInfo.Actions, actInfo)
			}
		}
		score, ok, err := s.markService.VerificationQuestionScore(&dto.Question{
			ExamID:     examID,
			StudentID:  studentID,
			QuestionID: res.QuestionID,
			Score:      res.Score})
		if err != nil {
			log.Error(err)
			return nil, false, err
		}
		if !ok {
			correct = false
		}
		qusInfo.Score = score.Score
		qusInfo.ScoredBy = score.ScoredBy
		qusInfo.ScoredTime = score.ScoredTime
		process.Questions = append(process.Questions, qusInfo)
	}

	return process, correct, nil
}
