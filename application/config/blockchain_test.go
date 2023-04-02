package config

import (
	"encoding/json"
	"github.com/liuhdd/exam-cret/application/models"
	"github.com/stretchr/testify/assert"
	"testing"
)



func Test_loadConfig(t *testing.T) {

	LoadConfig()
	assert.NotEmpty(t, c)
}

func TestGetContract(t *testing.T) {

	contract := GetContract()
	res, err := contract.SubmitTransaction("QueryActionByID", "action1")
	if err != nil {
		t.Fatal(err)
	}
	a := &models.ExamAction{}
	json.Unmarshal(res, a)
	t.Log(a)
}
