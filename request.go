package go_flow_task_control

type Request interface {
	IsStopped() bool
	IsFailed() bool
	Fail()
	Stop()
}

func NewRequest() Request {
	return &request{
		response:   NewResponse(),
		requestMap: map[string]interface{}{},
	}
}

type request struct {
	response   Response
	requestMap map[string]interface{}
}

func (req *request) IsFailed() bool {
	return req.response.Failed()
}

func (req *request) Fail() {
	req.response.SetStopped(true)
	req.response.SetFailed(true)
}

func (req *request) Stop() {
	req.response.SetStopped(true)
}

func (req *request) IsStopped() bool {
	return req.response.Stopped()
}
