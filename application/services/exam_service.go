package services

import (
	"github.com/liuhdd/exam-cret/application/config"
	"github.com/liuhdd/exam-cret/application/models"
	"github.com/liuhdd/exam-cret/application/repository"
	"github.com/liuhdd/exam-cret/application/services/dto"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
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
	FindExamRecordByExamIDAndStudentID(examID string, studentID string) (*models.ExamRecord, error)
	SaveQuestion(question *models.Question) error
	FindQuestionByID(id string) (*models.Question, error)
}

type examService struct {
	examRepo      repository.ExamRepository
	actionService ActionService
	markService   MarkService
	db			*gorm.DB
}

func NewExamService() ExamService {
	return &examService{
		actionService: NewActionService(),
		markService:   NewMarkService(),
		db: config.GetDB(),
	}
}

func (s *examService) FindQuestionByID(id string) (*models.Question, error) {
	var question models.Question
	tx := s.db.Where("question_id=?", id).First(&question)
	if tx.Error != nil {
		log.Error(tx.Error)
		return nil, tx.Error
	}
	return &question, nil
}

func (s *examService) SaveQuestion(question *models.Question) error {
	tx := s.db.Save(question)
	if tx.Error != nil {
		log.Error(tx.Error)
		return tx.Error
	}
	return nil
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

func (s *examService) FindExamRecordByExamIDAndStudentID(examID string, studentID string) (*models.ExamRecord, error) {
	return s.examRepo.GetExamRecordByExamIDAndStudentID(examID, studentID)
}

func (s *examService) FindExamsByStudentID(id string) ([]*models.Exam, error) {
	return s.examRepo.GetExamsByStudentID(id)
}
func (s *examService) FindExamResultByExamIDAndStudentID(examID, studentID string) (*dto.ExamResult, error) {
	
	var result []*dto.QuestionResult
	tx := s.db.Raw("select mark.question_id as question_id, content, answer, mark.score as score" +
	"from mark left join question on mark.question_id = question.question_id" +
	"where exam_id = ? and student_id = ?", examID, studentID).
	Scan(&result)
	if tx.Error != nil {
		return nil, tx.Error
	}

	var exam models.Exam
	tx = s.db.Where("exam_id=?", examID).First(&exam)
	if tx.Error != nil {
		return nil, tx.Error
	}

	var examRecord models.ExamRecord
	tx = s.db.Where("exam_id=? and student_id=?", examID, studentID).First(&examRecord)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &dto.ExamResult{
		ExamID:    examID,
		StudentID: studentID,
		ExamName: exam.ExamName,
		BeginTime: exam.BeginTime,
		EndTime: exam.EndTime,
		Grade: examRecord.Grade,
		Score: examRecord.Score,
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
		if score != nil {
			qusInfo.Score = score.Score
			qusInfo.ScoredBy = score.ScoredBy
			qusInfo.ScoredTime = score.ScoredTime

		}
		process.Questions = append(process.Questions, qusInfo)
	}

	return process, correct, nil
}
