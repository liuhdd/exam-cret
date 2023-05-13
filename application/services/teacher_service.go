package services

import (
	"github.com/liuhdd/exam-cret/application/config"
	"github.com/liuhdd/exam-cret/application/models"
	"gorm.io/gorm"
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
	db *gorm.DB
}
var t *teacherService

func NewTeacherService() TeacherService {
	once.Do(func() {
		db := config.GetDB()
		db.AutoMigrate(&models.Teacher{})
		t = &teacherService{db: db}
	})
	return t
}

func (t *teacherService) CreateTeacher(teacher *models.Teacher) error {
	tx := t.db.Save(teacher)
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
	tx := t.db.Save(teacher)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (t *teacherService) DeleteTeacher(id string) error {
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

