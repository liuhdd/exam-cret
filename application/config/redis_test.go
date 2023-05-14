package config

import (
	"context"
	"github.com/liuhdd/exam-cret/application/models"
	"testing"
)

func TestRedis(t *testing.T) {
	rdb := GetRedisClient()
	stu := &models.Student{
		StudentID: "123123",
		Name:      "刘",
	}
	ctx := context.Background()
	rdb.HSet(ctx, stu.StudentID, stu)

}
