package main

import (
	"testing"
)

func TestExamAction_UnmarshalJSON(t *testing.T) {

}

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}
