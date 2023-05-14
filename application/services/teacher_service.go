package services

import (
	"github.com/liuhdd/exam-cret/application/config"
	"github.com/liuhdd/exam-cret/application/models"
	"gorm.io/gorm"
	"sync"
)

type TeacherService interface {
	CreateTeacher(teacher *models.Teacher) error
	GetTeacherByID(id string) (*models.Teacher, error)
	GetTeacherByName(name string) ([]*models.Teacher, error)
	GetAllTeachers() ([]*models.Teacher, error)
	UpdateTeacher(teacher *models.Teacher) error
	DeleteTeacher(id string) error
}

type teacherService struct {
	db          *gorm.DB
	authService AuthService
}

var t *teacherService
var to sync.Once

func NewTeacherService() TeacherService {
	to.Do(func() {
		db := config.GetDB()
		db.AutoMigrate(&models.Teacher{})
		t = &teacherService{
			db:          db,
			authService: NewAuthService(),
		}
	})
	return t
}

func (t *teacherService) CreateTeacher(teacher *models.Teacher) error {

	teacher.UserID = teacher.TeacherID
	teacher.Username = teacher.TeacherID
	teacher.Password = "123456"
	teacher.Role = "teacher"
	u := models.User{
		UserID:   teacher.UserID,
		Username: teacher.Username,
		Password: teacher.Password,
		Role:     "teacher",
	}

	err := t.authService.Register(&u)
	if err != nil {
		return err
	}
	tx := t.db.Create(teacher)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (t *teacherService) GetTeacherByID(id string) (*models.Teacher, error) {
	var teacher models.Teacher
	tx := t.db.Where("teacher_id=?", id).First(&teacher)

	if tx.Error != nil {
		return nil, tx.Error
	}
	return &teacher, nil
}

func (t *teacherService) GetAllTeachers() ([]*models.Teacher, error) {
	var teachers []*models.Teacher
	tx := t.db.Find(&teachers)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return teachers, nil
}

func (t *teacherService) UpdateTeacher(teacher *models.Teacher) error {
	id := teacher.TeacherID
	if id == "" {
		return gorm.ErrInvalidData
	}
	var te *models.Teacher
	t.db.Where("teacher_id = ?", id).Find(&te)
	if te == nil {
		return gorm.ErrRecordNotFound
	}
	tx := t.db.Save(teacher)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (t *teacherService) DeleteTeacher(id string) error {
	t.db.Delete(&models.User{}, id)
	tx := t.db.Where("teacher_id=?", id).Delete(&models.Teacher{})
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (t *teacherService) GetTeacherByName(name string) ([]*models.Teacher, error) {
	var teachers []*models.Teacher
	tx := t.db.Where("name=?", name).Find(&teachers)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return teachers, nil
}
