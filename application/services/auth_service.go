package services

import (
	"context"
	"fmt"
	"github.com/liuhdd/exam-cret/application/config"
	"github.com/liuhdd/exam-cret/application/models"
	"github.com/liuhdd/exam-cret/application/repository"
	"github.com/redis/go-redis/v9"
	"golang.org/x/crypto/bcrypt"
	"log"
	"sync"
)

type AuthService interface {
	Register(user *models.User) error
	Login(user *models.User) error
}

var auth *authService

type authService struct {
	repo repository.UserRepository
	rdb  *redis.Client
}

var ao sync.Once

func NewAuthService() AuthService {
	ao.Do(func() {
		auth = &authService{
			repo: repository.NewUserRepository(),
			rdb:  config.GetRedisClient(),
		}
	})

	return auth
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
	if user.Role == "student" {
		ss := NewStudentService()
		stu, err := ss.GetStudentByID(user.UserID)
		if err != nil {
			return err
		}

		err = as.rdb.HSet(context.Background(), stu.UserID, stu).Err()
		if err != nil {
			return err
		}
	}
	return nil
}
