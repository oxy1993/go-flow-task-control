package go_flow_task_control

type Task interface {
	Run(T request, R response)
	RunAsync(T request, R response)
}

type task struct {
}

func (t *task) Execute(T request, R response) {}

func (t *task) Run(T request, R response) {
	t.Execute(T, R)
}

func (t *task) RunAsync(T request, R response) {
	go t.Execute(T, R)
}
