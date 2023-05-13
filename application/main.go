package main

import (
	"github.com/liuhdd/exam-cret/application/jobs"
	"github.com/liuhdd/exam-cret/application/routes"
)

func main() {

	engine := routes.InitEngine()
	jobs.ActionJob()
	engine.Run()
}
