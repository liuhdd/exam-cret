package repository

import (
	"github.com/liuhdd/exam-cret/application/config"
	"github.com/liuhdd/exam-cret/application/models"
	"gorm.io/gorm"
)

type StudentRepository interface {
	// GetStudentByID gets a student by ID
	GetStudentByID(id string) (*models.Student, error)
	SaveStudent(student *models.Student) error
	DeleteStudent(id string) error
}

type studentRepository struct {
	StudentRepository
	db *gorm.DB
}

func NewStudentRepository() StudentRepository {
	db := config.GetDB()
	db.AutoMigrate(&models.Student{})
	repo := &studentRepository{db: db}
	return repo
}

func (r *studentRepository) GetStudentByID(id string) (*models.Student, error) {
	var student models.Student
	tx := r.db.Where("student_id=?", id).First(&student)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &student, nil
}

func (r *studentRepository) SaveStudent(student *models.Student) error {
	tx := r.db.Save(student)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (r *studentRepository) DeleteStudent(id string) error {
	tx := r.db.Where("student_id=?", id).Delete(&models.Student{})
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}