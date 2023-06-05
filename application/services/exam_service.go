package services

import (
	"context"
	"github.com/gin-gonic/gin"
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
	DeleteExam(id string) error
	GetGradesByStudentID(studentID string) ([]*dto.Grade, error)
	ListExams() []*models.Exam
	QueryExam(c *gin.Context, exam *models.Exam) []*models.Exam
	QueryGrades(exam *models.ExamRecord) []*dto.Grade
	CreateQuestions(questions []*models.Question) error
}

type examService struct {
	examRepo      repository.ExamRepository
	actionService ActionService
	markService   MarkService
	db            *gorm.DB
}

func (s *examService) CreateQuestions(questions []*models.Question) error {
	tx := s.db.Save(questions)
	return tx.Error
}

func NewExamService() ExamService {
	return &examService{
		examRepo:      repository.NewExamRepository(),
		actionService: NewActionService(),
		markService:   NewMarkService(),
		db:            config.GetDB(),
	}
}

func (s *examService) QueryGrades(exam *models.ExamRecord) []*dto.Grade {
	var grades []*dto.Grade

	sub1 := s.db.Table("exam_records").Where(exam).Select("*")

	s.db.Table("(?) as er, students s, exams e", sub1).Where("er.student_id = s.student_id and er.exam_id = e.exam_id").
		Select("er.exam_id as exam_id, e.exam_name as exam_name, er.student_id as student_id, " +
			"s.name as student_name,  er.grade as grade").Scan(&grades)
	return grades
}
func (s *examService) ListExams() []*models.Exam {
	var exams []*models.Exam
	s.db.Find(&exams)
	return exams
}

func (s *examService) GetGradesByStudentID(studentID string) ([]*dto.Grade, error) {
	var grades []*dto.Grade
	tx := s.db.Raw("select er.exam_id as exam_id, e.exam_name as exam_name, er.student_id as student_id,"+
		" s.name as student_name, er.grade as grade"+
		" from exam_records er, students s, exams e "+
		"where er.student_id = ? and er.student_id = s.student_id and er.exam_id = e.exam_id",
		studentID).Scan(&grades)

	if tx.Error != nil {
		return nil, tx.Error
	}
	return grades, nil
}

func (s *examService) QueryExam(c *gin.Context, exam *models.Exam) (exams []*models.Exam) {

	get, exists := c.Get("user")
	if exists {
		if user, ok := get.(models.User); ok {
			if user.Role == "student" {
				s.db.Raw("select * from exams where exam_id in "+
					"(select exam_id from exam_records where student_id = ?)", user.Username).Scan(exams)
				return
			}
		}

	}
	s.db.Table("exams").Where(exam).Scan(&exams)
	return
}

func (s *examService) DeleteExam(id string) error {
	tx := s.db.Where("exam_id = ?", id).Delete(&models.Exam{})
	if tx.Error != nil {
		return tx.Error
	}
	return nil
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
	rdb := config.GetRedisClient()
	var e *dto.ExamResult
	err := rdb.HGetAll(context.Background(), examID+":"+studentID).Scan(&e)
	if err != nil {
		log.Error(err)
	}
	if e != nil {
		return e, nil
	}

	var result []*dto.QuestionResult
	tx := s.db.Raw("select marks.question_id as question_id,question_type, content, options, marks.answer as answer, marks.score as score "+
		"from marks left join questions on marks.question_id = questions.question_id "+
		"where exam_id = ? and student_id = ?", examID, studentID).
		Scan(&result)
	name := ""
	tx.Table("students").Select("name").Where("student_id = ?", studentID).Scan(&name)
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
		ExamID:      examID,
		StudentID:   studentID,
		StudentName: name,
		ExamName:    exam.ExamName,
		BeginTime:   exam.BeginTime,
		EndTime:     exam.EndTime,
		Grade:       examRecord.Grade,
		Questions:   result,
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
