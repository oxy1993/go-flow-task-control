package gftc

type Task interface {
	Run(T Request, R Response)
	RunAsync(T Request, R Response, C chan bool)
}

type task struct {
	exec Execute
}

func NewTask(exec Execute) Task {
	return &task{exec}
}

type Execute func(T Request, R Response)

func (t *task) Run(T Request, R Response) {
	t.exec(T, R)
}

func (t *task) RunAsync(T Request, R Response, C chan bool) {
	t.exec(T, R)
	C <- true
	close(C)
}
