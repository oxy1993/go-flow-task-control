package main

import (
	"log"
	"oxy/example/gftc"
)

func main() {
	newFlow := gftc.NewFlow()

	task1 := gftc.NewTask(HelloTaskExecute)
	helloTask := HelloTask{}
	helloTask.Task = task1

	task2 := gftc.NewTask(GoodbyeTaskExecute)
	goodbyeTask := GoodbyeTask{}
	goodbyeTask.Task = task2

	newFlow.AddTask(helloTask, false, false)
	newFlow.AddTask(goodbyeTask, false, false)

	request := gftc.NewRequest()
	response := gftc.NewResponse()

	newFlow.Run(request, response)
}

type HelloTask struct {
	gftc.Task
}

func HelloTaskExecute(t gftc.Request, r gftc.Response) {
	log.Println("HelloTask ===> Execute")
}

type GoodbyeTask struct {
	gftc.Task
}

func GoodbyeTaskExecute(t gftc.Request, r gftc.Response) {
	log.Println("GoodbyeTask ===> Execute")
}