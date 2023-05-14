package services

import (
	"errors"
	"github.com/liuhdd/exam-cret/application/config"
	"github.com/liuhdd/exam-cret/application/models"
	"github.com/liuhdd/exam-cret/application/repository"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"sync"
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
	CreateStudent(student *models.Student) error
}

type studentService struct {
	db          *gorm.DB
	authService AuthService
	studentRepo repository.StudentRepository
}

var s *studentService
var so sync.Once

func NewStudentService() StudentService {
	if s == nil {
		so.Do(func() {
			s = &studentService{}
			s.authService = NewAuthService()
			s.db = config.GetDB()
			s.studentRepo = repository.NewStudentRepository()

		})
	}
	return s
}

func (s *studentService) CreateStudent(student *models.Student) error {
	student.Password = "123456"
	student.UserID = student.StudentID
	student.Username = student.StudentID
	student.Role = "student"
	u := models.User{
		UserID:   student.UserID,
		Username: student.Username,
		Password: student.Password,
		Role:     "student",
	}
	err := s.authService.Register(&u)
	if err != nil {
		return err
	}
	tx := s.db.Create(student)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
func (s *studentService) UpdateStudent(student *models.Student) error {
	id := student.StudentID
	if id == "" {
		return errors.New("missing id")
	}
	var stu *models.Student
	s.db.Where("id = ?", id).Scan(stu)
	if stu == nil {
		return gorm.ErrRecordNotFound
	}
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
	tx := s.db.Where("name=?", name).Find(&students)
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
	s.db.Delete(&models.User{}, id)
	return s.studentRepo.DeleteStudent(id)
}
