package config_test

import (
	"fmt"
	"testing"

	"github.com/liuhdd/exam-cret/application/config"
	"github.com/stretchr/testify/assert"
)

func TestGetRootPath(t *testing.T) {
	path := config.GetRootPath()
	fmt.Println(123123)
	t.Log(path)
}

func TestLoadConfig(t *testing.T) {
	config.LoadConfig()
	assert.Equal(t, "debug", config.GetProperty("environment"))
}