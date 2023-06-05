package jobs

import (
	"context"
	"github.com/liuhdd/exam-cret/application/repository"
	"strconv"

	"github.com/liuhdd/exam-cret/application/config"
	"github.com/liuhdd/exam-cret/application/models"
	"github.com/redis/go-redis/v9"
	log "github.com/sirupsen/logrus"
)

func Setup() {
	ActionJob()
	MarkJob()
}
func ActionJob() {
	rdb := config.GetRedisClient()
	args := &redis.XReadGroupArgs{
		Group:    "g1",
		Consumer: "upload",
		Streams:  []string{"action", ">"},
		Count:    100,
		Block:    0,
	}
	ctx := context.Background()
	go func() {
		for true {
			res, err := rdb.XReadGroup(ctx, args).Result()
			if err != nil {
				log.Errorf("redis xread error: %v", err)
			}
			upload(ctx, rdb, res)
		}
	}()

}

func upload(ctx context.Context, rdb *redis.Client, res []redis.XStream) {
	repo := repository.NewActionRepository()
	for _, stream := range res {
		for _, message := range stream.Messages {
			time, _ := strconv.ParseInt(message.Values["answer_time"].(string), 10, 0)
			t, _ := strconv.ParseUint(message.Values["action_type"].(string), 10, 0)
			action := &models.ExamAction{
				ObjectType: message.Values["object_type"].(string),
				ActionID:   message.Values["action_id"].(string),
				ExamID:     message.Values["exam_id"].(string),
				StudentID:  message.Values["student_id"].(string),
				QuestionID: message.Values["question_id"].(string),
				Answer:     message.Values["answer"].(string),
				ActionTime: time,
				ActionType: uint(t),
			}

			err := repo.AddAction(action)
			if err != nil {
				log.Errorf("upload action error: %v", err)
			}
			rdb.XAck(ctx, "action", "action", message.ID).Result()
		}
	}
}
