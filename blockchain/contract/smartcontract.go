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
func (sc *ActionContract) QueryAction(ctx contractapi.TransactionContextInterface) {

}
