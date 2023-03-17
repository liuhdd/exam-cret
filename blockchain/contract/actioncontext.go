package main

import (
	"encoding/json"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type ActionContextInterface interface {
	contractapi.TransactionContextInterface
	AddAction(action *ExamAction) error
	QueryActionByID(actionID string) (*ExamAction, error)
}

type ActionContext struct {
	contractapi.TransactionContext
}

func (ctx *ActionContext) AddAction(action *ExamAction) error {
	data, err := json.Marshal(action)
	if err != nil {
		return err
	}
	err = ctx.GetStub().PutState(action.ActionID, data)
	if err != nil {
		return err
	}
	return nil
}

func (ctx *ActionContext) QueryActionByID(actionID string) (*ExamAction, error) {
	state, err := ctx.GetStub().GetState(actionID)
	if err != nil {
		return nil, err
	}
	action := &ExamAction{}
	json.Unmarshal(state, action)
	return action, nil
}
