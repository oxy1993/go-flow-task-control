package go_flow_task_control

import (
	"golang.org/x/exp/slices"
	"log"
)

type Flow interface {
	Run(T request, R response)
}

type flow struct {
	tasks     []Task
	alwaysRun []int
	asyncRun  []int
}

func (f *flow) Run(t request, r response) {
	t.response = &r
	f.runATask(t, r, 0)
}

func (f *flow) runATask(T request, R response, index int) {
	if index == len(f.tasks) {
		return
	}

	if T.IsStopped() || slices.Contains(f.alwaysRun, index) {
		if slices.Contains(f.asyncRun, index) {
			log.Printf("Run async task %s at index %d", f.tasks[index], index)
			go f.tasks[index].Run(T, R)
		} else {
			f.tasks[index].Run(T, R)
		}
	}
}

func (f *flow) AddTask(task Task, isAlwaysRun bool, isRunAsync bool) {
	f.tasks = append(f.tasks, task)

	if isAlwaysRun {
		_ = append(f.alwaysRun, len(f.tasks)-1)
	}

	if isRunAsync {
		_ = append(f.asyncRun, len(f.tasks)-1)
	}
}
