package gftc

type IRequest interface {
	IsStopped() bool
	IsFailed() bool
	Fail()
	Stop()
}

func NewRequest() Request {
	return Request{
		Response:   NewResponse(),
		RequestMap: map[string]interface{}{},
	}
}

type Request struct {
	Response   Response
	RequestMap map[string]interface{}
}

func (req *Request) IsFailed() bool {
	return req.Response.Failed()
}

func (req *Request) Fail() {
	req.Response.SetStopped(true)
	req.Response.SetFailed(true)
}

func (req *Request) Stop() {
	req.Response.SetStopped(true)
}

func (req *Request) IsStopped() bool {
	return req.Response.Stopped()
}
