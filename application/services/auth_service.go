package services

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/liuhdd/exam-cret/application/config"
	"github.com/liuhdd/exam-cret/application/models"
	"github.com/liuhdd/exam-cret/application/repository"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"log"
	"sync"
)

type AuthService interface {
	Register(user *models.User) error
	Login(user *models.User) (string, error)
	ListUsers() []*models.User
	DeleteUser(username string)
}

var auth *authService

type authService struct {
	repo repository.UserRepository
	rdb  *redis.Client
	db   *gorm.DB
}

func (as *authService) DeleteUser(username string) {
	as.db.Where("username = ?", username).Delete(&models.User{})
}

var ao sync.Once

func NewAuthService() AuthService {
	ao.Do(func() {
		auth = &authService{
			repo: repository.NewUserRepository(),
			rdb:  config.GetRedisClient(),
			db:   config.GetDB(),
		}
	})

	return auth
}
func (as *authService) ListUsers() []*models.User {
	var users []*models.User
	as.db.Find(&users)
	return users
}
func (as *authService) Register(user *models.User) error {

	err := as.repo.CreateUser(user)
	if err != nil {
		log.Printf("Failed to create user: %v", err)
		return err
	}
	return nil
}

func (as *authService) Login(user *models.User) (string, error) {

	u, err := as.repo.FindUserByUserName(user.Username)
	if err != nil {
		log.Printf("Failed to query user by username: %s", err)
		return "", err
	}
	if u.Password != user.Password {
		return "", fmt.Errorf("username or password wrong")
	}
	user.Role = u.Role
	token := uuid.New().String()
	if u.Role == "student" {
		ss := NewStudentService()
		stu, err := ss.GetStudentByID(user.Username)
		if err != nil {
			return "", err
		}
		ctx := context.Background()
		err = as.rdb.HSet(ctx, token, stu).Err()
		if err != nil {
			return "", err
		}

		examSer := NewExamService()
		exams, _ := examSer.FindExamsByStudentID(stu.StudentID)
		if err != nil {
			for _, e := range exams {
				res, _ := examSer.FindExamResultByExamIDAndStudentID(e.ExamID, stu.StudentID)
				err := as.rdb.HSet(ctx, e.ExamID+":"+stu.StudentID, res).Err()
				if err != nil {
					return "", err
				}
			}

		}

	}

	return token, nil
}
