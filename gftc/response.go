package gftc

import "time"

type IResponse interface {
	GetResponseMap() map[string]interface{}
	SetResponseMap(mapResponse map[string]interface{})
	Stopped() bool
	SetStopped(stopped bool)
	Failed() bool
	SetFailed(failed bool)
	Result() int32
	SetResult(result int32)
}

type Response struct {
	time        time.Time
	stopped     bool
	failed      bool
	result      int32
	responseMap map[string]interface{}
}

func NewResponse() Response {
	return Response{time: time.Now()}
}

func (res *Response) Stopped() bool {
	return res.stopped
}

func (res *Response) SetStopped(stopped bool) {
	res.stopped = stopped
}

func (res *Response) Failed() bool {
	return res.failed
}

func (res *Response) SetFailed(failed bool) {
	res.failed = failed
}

func (res *Response) Result() int32 {
	return res.result
}

func (res *Response) SetResult(result int32) {
	res.result = result
}

func (res *Response) GetResponseMap() map[string]interface{} {
	return res.responseMap
}

func (res *Response) SetResponseMap(responseMap map[string]interface{}) {
	res.responseMap = responseMap
}
