package services

import (
	"github.com/liuhdd/exam-cret/application/config"
	"github.com/liuhdd/exam-cret/application/models"
	"github.com/liuhdd/exam-cret/application/repository"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type StudentService interface {
	// GetStudentByID gets a student by ID
	GetStudentByID(id string) (*models.Student, error)
	SaveStudent(student *models.Student) error
	DeleteStudent(id string) error
	GetStudentsByExamId(id string) ([]*models.Student, error)
	GetStudentByName(name string) ([]*models.Student, error)
	GetAllStudents() ([]*models.Student, error)
	UpdateStudent(student *models.Student) error
}

type studentService struct {
	db *gorm.DB
	studentRepo repository.StudentRepository
}

var s *studentService
func NewStudentService() StudentService {
	once.Do(func() {
		db := config.GetDB()
		s = &studentService{db: db, studentRepo: repository.NewStudentRepository()}
	})
	return s
}

func (s *studentService) UpdateStudent(student *models.Student) error {
	tx := s.db.Save(student)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (s *studentService) GetAllStudents() ([]*models.Student, error) {
	var students []*models.Student
	tx := s.db.Find(&students)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return students, nil
}

func (s *studentService) GetStudentByName(name string) ([]*models.Student, error) {
	var students []*models.Student
	tx := s.db.Where("student_name=?", name).Find(&students)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return students, nil
}

func (s *studentService) GetStudentsByExamId(id string) ([]*models.Student, error) {
	var students []*models.Student
	tx := s.db.Raw("select * from student where student_id in (select student_id from exam_record where exam_id = ?)", id).Scan(&students)
	log.Error(tx.Error)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return students, nil

	
}

func (s *studentService) GetStudentByID(id string) (*models.Student, error) {
	return s.studentRepo.GetStudentByID(id)
}

func (s *studentService) SaveStudent(student *models.Student) error {
	return s.studentRepo.SaveStudent(student)
}

func (s *studentService) DeleteStudent(id string) error {
	return s.studentRepo.DeleteStudent(id)
}
