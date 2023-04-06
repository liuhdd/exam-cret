package main

import (
	"fmt"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type ActionContract struct {
	contractapi.Contract
}

func (sc *ActionContract) UploadAction(ctx ActionContextInterface,
	actionID string, examID string, studentID string, actionType uint,
	time int64, questionID string, answer string) error {

	json, err := ctx.GetStub().GetState(actionID)
	if err != nil && len(json) > 0 {
		return fmt.Errorf("action: %s existed", actionID)
	}

	action := &ExamAction{
		ObjectType: "exam_action",
		ActionID:   actionID,
		ExamID:     examID,
		StudentID:  studentID,
		ActionType: actionType,
		ActionTime: time,
		QuestionID: questionID,
		Answer:     answer,
	}
	err = ctx.AddAction(action)
	if err != nil {
		return fmt.Errorf("failed to uploead exam action: %s, with error: %s", examID, err)
	}
	return nil
}
func (sc *ActionContract) QueryActionByID(ctx ActionContextInterface,
	id string) (*ExamAction, error) {
	action, err := ctx.QueryActionByID(id)
	if err != nil {
		return nil, fmt.Errorf("failed to query exam action: %s, with error: %s", id, err)
	}
	if action == nil {
		return nil, fmt.Errorf("action not found: %s", id)
	}

	return action, nil

}

func (sc *ActionContract) QueryActionByExamAndStudentID(ctx ActionContextInterface, objectType, examID, studentID string) (string, error) {
	actions, err := ctx.GetActionByExamAndStudentID(objectType, examID, studentID)
	if err != nil {
		return "", fmt.Errorf("error with querying actions: %s", err)
	}
	return actions, nil

}
func (sc *ActionContract) QueryAction(ctx ActionContextInterface, queryJson string) ([]*ExamAction, error) {
	return ctx.QueryAction(queryJson)
}

func (sc *ActionContract) UploadMarkAction(ctx ActionContextInterface,
	actionID string, examID string, studentID string, questionID string,
	score uint, time int64, scoreBy string) error {

	json, err := ctx.GetStub().GetState(actionID)
	if err != nil && len(json) > 0 {
		return fmt.Errorf("action: %s existed", actionID)
	}

	action := &MarkAction{
		ObjectType: "mark_action",
		ActionID:   actionID,
		ExamID:     examID,
		StudentID:  studentID,
		QuestionID: questionID,
		Scorer:     scoreBy,
		Score:      score,
		ScoredTime: time,
	}
	err = ctx.AddMarkAction(action)
	if err != nil {
		return fmt.Errorf("failed to uploead exam action: %s, with error: %s", examID, err)
	}
	return nil
}

func (sc *ActionContract) QueryMarkActionByID(ctx ActionContextInterface, actionID string) (*MarkAction, error) {
	action, err := ctx.QueryMarkActionByID(actionID)
	if err != nil {
		return nil, fmt.Errorf("failed to query exam action: %s, with error: %s", actionID, err)
	}
	if action == nil {
		return nil, fmt.Errorf("action not found: %s", actionID)
	}

	return action, nil

}

func (sc *ActionContext) QuestionScore(ctx ActionContextInterface, examID, studentID, questionID string) (int, error) {
	return ctx.GetQuestionScore(examID, studentID, questionID)
}
