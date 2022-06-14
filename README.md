# go-flow-task-control

Flow-Task-Control is a library for clear organization of code and easy code reuse provided by Oxy.

We divide the work into multiple Tasks and attach it to the respective Flows. Each Flow will contain multiple tasks according to their business logic. A Task can be used for many Flows.

Flow is a Functional Interface with only a run method. After adding a task into the flow. We just need to call the run function to start the Flow.

When adding a Task to the Flow, we can specify whether the Task should always be run or not and whether to run async. We can also run a task inside a task, specifically in the Usage section. Now, let’s move on to The Installation.

$ go get github.com/oxy1993/go-flow-task-control