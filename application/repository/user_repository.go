package repository

import (
	"github.com/liuhdd/exam-cret/application/config"
	"github.com/liuhdd/exam-cret/application/models"
	"gorm.io/gorm"
)

func GetDB() *gorm.DB {
	return config.GetDB()
}

type UserRepository interface {
	CreateUser(user *models.User) error
	FindUserByUserName(username string) (*models.User, error)
}

func NewUserRepository() UserRepository {
	db := config.GetDB()
	db.AutoMigrate(&models.User{})
	repo := &userRepository{db}
	return repo
}

type userRepository struct {
	db *gorm.DB
}

func (r *userRepository) CreateUser(user *models.User) error {
	tx := r.db.Create(user)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (r *userRepository) FindUserByUserName(username string) (*models.User, error) {
	var user models.User
	tx := r.db.Where("username=?", username).First(&user)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &user, nil
}
