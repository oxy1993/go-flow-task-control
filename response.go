package go_flow_task_control

import "time"

type Response interface {
	GetResponseMap() map[string]interface{}
	SetResponseMap(mapResponse map[string]interface{})
	Stopped() bool
	SetStopped(stopped bool)
	Failed() bool
	SetFailed(failed bool)
	Result() int32
	SetResult(result int32)
}

type response struct {
	time        time.Time
	stopped     bool
	failed      bool
	result      int32
	responseMap map[string]interface{}
}

func NewResponse() Response {
	return &response{time: time.Now()}
}

func (res *response) Stopped() bool {
	return res.stopped
}

func (res *response) SetStopped(stopped bool) {
	res.stopped = stopped
}

func (res *response) Failed() bool {
	return res.failed
}

func (res *response) SetFailed(failed bool) {
	res.failed = failed
}

func (res *response) Result() int32 {
	return res.result
}

func (res *response) SetResult(result int32) {
	res.result = result
}

func (res *response) GetResponseMap() map[string]interface{} {
	return res.responseMap
}

func (res *response) SetResponseMap(responseMap map[string]interface{}) {
	res.responseMap = responseMap
}
