package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

var selectors map[string]string

func init() {
	selectors = make(map[string]string)
	selectors["GetActionByExamAndStudentID"] = "{\n  \"selector\": {\n    \"object_type\": \"%s\",\n   " +
		" \"exam_id\": \"%s\",\n    \"student_id\": \"%s\"\n  }}"
	selectors["GetQuestionScore"] = "{\n  \"selector\": {\n    \"object_type\": \"mark_action\",\n    \"exam_id\": \"%s\",\n    " +
		"\"student_id\": \"%s\",\n    \"question_id\": \"%s\"\n  }}"

}

type ActionContextInterface interface {
	contractapi.TransactionContextInterface
	AddAction(action *ExamAction) error
	QueryActionByID(actionID string) (*ExamAction, error)
	GetActionByExamAndStudentID(string, string, string) (string, error)
	QueryAction(query string) ([]*ExamAction, error)
	AddMarkAction(action *MarkAction) error
	QueryMarkActionByID(actionID string) (*MarkAction, error)
	GetQuestionScore(examID, studentID, questionID string) (int, error)
}

type ActionContext struct {
	contractapi.TransactionContext
}

func (ctx *ActionContext) GetQuestionScore(examID, studentID, questionID string) (int, error) {
	res, err := getQueryResultForQueryString(ctx.GetStub(), fmt.Sprintf(selectors["GetQuestionScore"], examID, studentID, questionID))
	if err != nil {
		return -1, err
	}
	score, err := strconv.Atoi(res)
	if err != nil {
		return -1, err
	}
	return score, nil
}

func (ctx *ActionContext) QueryMarkActionByID(actionID string) (*MarkAction, error) {
	state, err := ctx.GetStub().GetState(actionID)
	if err != nil {
		return nil, err
	}
	action := &MarkAction{}
	err = json.Unmarshal(state, action)
	if err != nil {
		return nil, err
	}
	return action, nil
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
	err = json.Unmarshal(state, action)
	if err != nil {
		return nil, err
	}
	return action, nil
}

func (ctx *ActionContext) AddMarkAction(action *MarkAction) error {
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

func (ctx *ActionContext) GetActionByExamAndStudentID(objectType, examID, studentID string) (string, error) {
	if objectType == "" || examID == "" || studentID == "" {
		return "", errors.New("args miss")
	}
	queryString := selectors["GetActionByExamAndStudentID"]
	actions, err := getQueryResultForQueryString(ctx.GetStub(), fmt.Sprintf(queryString, objectType, examID, studentID))
	if err != nil {
		return "", err
	}
	return actions, nil
}

func (ctx *ActionContext) QueryAction(queryJson string) ([]*ExamAction, error) {
	if queryJson == "" {
		return nil, errors.New("argument query json miss")
	}
	result, err := getQueryResultForQueryString(ctx.GetStub(), queryJson)
	if err != nil {
		return nil, err
	}
	var actions []*ExamAction
	err = json.Unmarshal([]byte(result), &actions)
	if err != nil {
		return nil, err
	}

	return actions, nil
}

func getQueryResultForQueryString(stub shim.ChaincodeStubInterface, queryString string) (string, error) {

	fmt.Printf("- getQueryResultForQueryString queryString:\n%s\n", queryString)

	resultsIterator, err := stub.GetQueryResult(queryString)
	if err != nil {
		return "", err
	}
	defer resultsIterator.Close()

	buffer, err := constructQueryResponseFromIterator(resultsIterator)
	if err != nil {
		return "", err
	}
	fmt.Printf("- getQueryResultForQueryString result:\n%s\n", buffer.String())
	return buffer.String(), nil
}

func constructQueryResponseFromIterator(resIterator shim.StateQueryIteratorInterface) (*bytes.Buffer, error) {
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resIterator.HasNext() {
		next, err := resIterator.Next()

		if err != nil {
			return nil, err
		}
		if bArrayMemberAlreadyWritten {
			buffer.WriteString(",")
		}
		buffer.Write(next.GetValue())
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")
	return &buffer, nil
}
