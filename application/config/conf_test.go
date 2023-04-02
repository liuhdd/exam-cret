package config_test

import (
	"fmt"
	"testing"

	"github.com/liuhdd/exam-cret/application/config"
)

func TestGetRootPath(t *testing.T) {
	path := config.GetRootPath()
	fmt.Println(123123)
	t.Log(path)
}