package main

import (
	"github.com/liuhdd/exam-cret/application/routes"
)

func main() {
	engine := routes.InitEngine()
	engine.Run()
}
