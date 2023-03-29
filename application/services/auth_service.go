package services

import (
	"fmt"
	"github.com/liuhdd/exam-cret/application/models"
	"github.com/liuhdd/exam-cret/application/repository"
	"golang.org/x/crypto/bcrypt"
	"log"
)

type AuthService interface {
	Register(user *models.User) error
	Login(user *models.User) error
}

type authService struct {
	repo repository.UserRepository
}

func NewAuthService(repo repository.UserRepository) AuthService {
	return &authService{repo}
}

func (as *authService) Register(user *models.User) error {
	hashedPasswd, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Failed to hash password: %s", err)
		return err
	}
	user.Password = string(hashedPasswd)
	err = as.repo.CreateUser(user)
	if err != nil {
		log.Printf("Failed to create user: %v", err)
		return err
	}
	return nil
}

func (as *authService) Login(user *models.User) error {

	u, err := as.repo.FindUserByUserName(user.Username)
	if err != nil {
		log.Printf("Failed to query user by username: %s", err)
		return err
	}
	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(user.Password))
	if err != nil {
		return fmt.Errorf("username or password wrong")
	}
	return nil
}
