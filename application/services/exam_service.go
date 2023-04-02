package services

import (

	"github.com/liuhdd/exam-cret/application/services/dto"
	log "github.com/sirupsen/logrus"
)

type ExamService interface {
	FindExamResultByExamIDAndStudentID(examID, studentID string) (*dto.ExamResult, error)
	VerifyExamResults(*dto.ExamResult) (*dto.ExamProcess, bool, error)
}

type examService struct {
	ExamService
	actionService ActionService
	markService MarkService
}

func NewExamService() ExamService {
	return &examService{
		actionService: NewActionService(),
		markService: NewMarkService(),
	}
}

func (s *examService) FindExamResultByExamIDAndStudentID(examID, studentID string) (*dto.ExamResult, error) {
	actions, err := as.actionRepo.GetAnswersFromDB(examID, studentID)
	if err != nil {
		log.Printf("error in find action: %s", err)
		return nil, err
	}
	var result []*dto.QuestionResult
	for _, action := range actions {
		result = append(result, &dto.QuestionResult{
			QuestionID: action.QuestionID,
			Answer:     action.Answer,
		})
	}
	return &dto.ExamResult{
		ExamID: examID,
		StudentID: studentID,
		Questions: result,
	}, nil
}

func (s *examService) VerifyExamResults(result *dto.ExamResult) (*dto.ExamProcess, bool, error) {
	examID := result.ExamID
	studentID := result.StudentID

	actions, err := s.actionService.SelectActionByExamAndStudentID(examID, studentID)
	if err != nil {
		log.Error(err)
		return nil, false, err
	}

	process := &dto.ExamProcess{
		ExamID: examID,
		StudentID: studentID,
	}

	var correct bool = true

	for _, res := range result.Questions {

		qusInfo := &dto.QuestionInfo{QuestionID: res.QuestionID}

		for _, action := range actions {

			if res.QuestionID == action.QuestionID {
				actInfo := &dto.ActionInfo{
					ActionID: action.ActionID,
					Answer: action.Answer,
					ActionTime: action.ActionTime,
				}
				qusInfo.Actions = append(qusInfo.Actions, actInfo)
			}
		}
		score, ok, err  := s.markService.VerificationQuestionScore(&dto.Question{
			ExamID: examID, 
			StudentID: studentID, 
			QuestionID: res.QuestionID, 
			Score: res.Score})
		if err != nil {
			log.Error(err)
			return nil, false, err
		}
		if !ok {
			correct = false
		}
		qusInfo.Score = score.Score
		qusInfo.ScoredBy = score.ScoredBy
		qusInfo.ScoredTime = score.ScoredTime
		process.Questions = append(process.Questions, qusInfo)
	}
	
	return process, correct, nil
}