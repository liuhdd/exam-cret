package jobs

import (
	"context"
	"github.com/liuhdd/exam-cret/application/config"
	"github.com/liuhdd/exam-cret/application/models"
	"github.com/liuhdd/exam-cret/application/repository"
	"github.com/redis/go-redis/v9"
	log "github.com/sirupsen/logrus"
	"strconv"
)

func MarkJob() {
	rdb := config.GetRedisClient()
	args := &redis.XReadGroupArgs{
		Group:    "g2",
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
			repo := repository.NewMarkRepository()
			for _, stream := range res {
				for _, message := range stream.Messages {
					time, _ := strconv.ParseInt(message.Values["answer_time"].(string), 10, 0)
					mark := &models.MarkAction{
						ObjectType: message.Values["object_type"].(string),
						ActionID:   message.Values["action_id"].(string),
						ExamID:     message.Values["exam_id"].(string),
						StudentID:  message.Values["student_id"].(string),
						QuestionID: message.Values["question_id"].(string),
						Score:      message.Values["score"].(uint),
						Scorer:     message.Values["scorer"].(string),
						ActionTime: time,
					}

					err := repo.UploadMarkToBC(mark)
					if err != nil {
						log.Errorf("upload action error: %v", err)
					}
					rdb.XAck(ctx, "action", "action", message.ID).Result()
				}
			}
		}
	}()
}
