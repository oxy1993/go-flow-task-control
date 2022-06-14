package main

import (
	"go-flow-task-control/gftc"
	"log"
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
	newFlow.AddTask(goodbyeTask, false, true)

	request := gftc.NewRequest()
	response := gftc.NewResponse()

	newFlow.Run(request, response)
}

type HelloTask struct {
	gftc.Task
}

func HelloTaskExecute(t gftc.Request, r gftc.Response) {
	t.RequestMap["oxy"] = 1993
	r.ResponseMap["mon"] = 1996
	log.Println("HelloTask ===> Execute")
}

type GoodbyeTask struct {
	gftc.Task
}

func GoodbyeTaskExecute(t gftc.Request, r gftc.Response) {
	log.Println("GoodbyeTask ===> Execute")
	log.Println("Oxy value in REQUEST added at HelloTask =>", t.RequestMap["oxy"])
	log.Println("Mon value in RESPONSE added at HelloTask =>", r.ResponseMap["mon"])
}
