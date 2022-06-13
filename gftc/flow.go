package gftc

import (
	"golang.org/x/exp/slices"
	"log"
)

type Flow interface {
	Run(T Request, R Response)
	AddTask(task Task, isAlwaysRun bool, isRunAsync bool)
}

type flow struct {
	tasks     []Task
	alwaysRun []int
	asyncRun  []int
}

func NewFlow() Flow {
	return &flow{
		tasks:     []Task{},
		alwaysRun: []int{},
		asyncRun:  []int{},
	}
}

func (f *flow) Run(t Request, r Response) {
	t.response = r
	f.runATask(t, r, 0)
}

func (f *flow) runATask(T Request, R Response, index int) {
	if index == len(f.tasks) {
		return
	}

	if !T.IsStopped() || slices.Contains(f.alwaysRun, index) {
		if slices.Contains(f.asyncRun, index) {
			log.Printf("Run async task %s at index %d", f.tasks[index], index)
			f.tasks[index].RunAsync(T, R)
		} else {
			f.tasks[index].Run(T, R)
		}
	}

	f.runATask(T, R, index+1)
}

func (f *flow) AddTask(task Task, isAlwaysRun bool, isRunAsync bool) {
	f.tasks = append(f.tasks, task)

	if isAlwaysRun {
		f.alwaysRun = append(f.alwaysRun, len(f.tasks)-1)
	}

	if isRunAsync {
		f.asyncRun = append(f.asyncRun, len(f.tasks)-1)
	}
}
