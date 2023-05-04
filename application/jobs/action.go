package jobs

import (
	"context"

	"github.com/liuhdd/exam-cret/application/config"
	"github.com/liuhdd/exam-cret/application/models"
	"github.com/liuhdd/exam-cret/application/repository"
	"github.com/redis/go-redis/v9"
	log "github.com/sirupsen/logrus"
)

func ActionJob() {
	rdb := config.GetRedisCilent()
	args := &redis.XReadArgs{
		Streams: []string{"action", "$"},
		Count:  0,
		Block:  0,
	}
	ctx := context.Background()
	res, err := rdb.XRead(ctx, args).Result()
	
	if err != nil {
		log.Errorf("redis xread error: %v", err)
	}

	repo := repository.NewActionRepository()
	for _, stream := range res {
		for _, message := range stream.Messages {
			action := &models.ExamAction{
				ObjectType: message.Values["object_type"].(string),
				ActionID:  message.Values["action_id"].(string),
				ExamID:     message.Values["exam_id"].(string),
				StudentID:  message.Values["student_id"].(string),
				QuestionID: message.Values["question_id"].(string),
				Answer:     message.Values["answer"].(string),
				ActionTime: message.Values["answer_time"].(int64),
				ActionType: message.Values["action_type"].(uint),
			}
			err = repo.AddAction(action)
			if err != nil {
				log.Errorf("upload action error: %v", err)
			}
			rdb.XAck(ctx, "action", "action", message.ID)
		}
	}
}