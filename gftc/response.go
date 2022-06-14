package gftc

import "time"

type IResponse interface {
}

type Response struct {
	Time        time.Time
	IsStopped   bool
	IsFailed    bool
	Result      int32
	ResponseMap map[string]interface{}
}

func NewResponse() Response {
	return Response{
		Time:        time.Now(),
		ResponseMap: map[string]interface{}{},
	}
}

func (res *Response) Failed() bool {
	return res.IsFailed
}

func (res *Response) Stopped() bool {
	return res.IsStopped
}

func (res *Response) SetStopped(stopped bool) {
	res.IsStopped = stopped
}

func (res *Response) SetFailed(failed bool) {
	res.IsFailed = failed
}
