package repository

import (
	"github.com/liuhdd/exam-cret/application/config"
	"github.com/liuhdd/exam-cret/application/models"
	"gorm.io/gorm"
)

type ExamRepository interface {
	
	GetExamByID(id string) (*models.Exam, error)
	
	GetExamsByStudentID(id string) ([]*models.Exam, error)
	
	GetExamsByName(name string) ([]*models.Exam, error)
	
	GetExamsByTime(beginTime int64, endTime int64) ([]*models.Exam, error)

	SaveExam(exam *models.Exam) error

	GetExamRecordByID(id string) (*models.ExamRecord, error)

	GetExamRecordsByStudentID(id string) ([]*models.ExamRecord, error)

	GetExamRecordsByExamID(id string) ([]*models.ExamRecord, error)

	SaveExamRecord(examRecord *models.ExamRecord) error

	GetExamRecordByExamIDAndStudentID(examID string, studentID string) (*models.ExamRecord, error)
}

type examRepository struct {
	
	db *gorm.DB
}

func NewExamRepository() ExamRepository {
	db := config.GetDB()
	db.AutoMigrate(&models.Exam{})
	db.AutoMigrate(&models.ExamRecord{})
	db.AutoMigrate(&models.Question{})
	repo := &examRepository{db: db}
	return repo
}

func (r *examRepository) GetExamByID(id string) (*models.Exam, error) {
	var exam models.Exam
	tx := r.db.Where("exam_id=?", id).First(&exam)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &exam, nil
}

func (r *examRepository) GetExamsByStudentID(id string) ([]*models.Exam, error) {
	var exams []*models.Exam
	tx := r.db.Where("student_id=?", id).Find(&exams)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return exams, nil
}

func (r *examRepository) GetExamsByName(name string) ([]*models.Exam, error) {
	var exams []*models.Exam
	tx := r.db.Where("name=?", name).Find(&exams)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return exams, nil
}

func (r *examRepository) GetExamsByTime(beginTime int64, endTime int64) ([]*models.Exam, error) {
	var exams []*models.Exam
	tx := r.db.Where("begin_time>=? and end_time<=?", beginTime, endTime).Find(&exams)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return exams, nil
}

func (r *examRepository) SaveExam(exam *models.Exam) error {
	tx := r.db.Save(exam)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (r *examRepository) GetExamRecordByID(id string) (*models.ExamRecord, error) {
	var examRecord models.ExamRecord
	tx := r.db.Where("exam_record_id=?", id).First(&examRecord)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &examRecord, nil
}

func (r *examRepository) GetExamRecordsByStudentID(id string) ([]*models.ExamRecord, error) {
	var examRecords []*models.ExamRecord
	tx := r.db.Where("student_id=?", id).Find(&examRecords)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return examRecords, nil
}

func (r *examRepository) GetExamRecordsByExamID(id string) ([]*models.ExamRecord, error) {
	var examRecords []*models.ExamRecord
	tx := r.db.Where("exam_id=?", id).Find(&examRecords)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return examRecords, nil
}

func (r *examRepository) SaveExamRecord(examRecord *models.ExamRecord) error {
	tx := r.db.Save(examRecord)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (r *examRepository) GetExamRecordByExamIDAndStudentID(examID string, studentID string) (*models.ExamRecord, error) {
	var examRecords *models.ExamRecord
	tx := r.db.Where("exam_id=? and student_id=?", examID, studentID).Find(&examRecords)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return examRecords, nil
}
