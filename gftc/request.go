package gftc

type IRequest interface {
	IsStopped() bool
	IsFailed() bool
	Fail()
	Stop()
}

func NewRequest() Request {
	return Request{
		response:   NewResponse(),
		requestMap: map[string]interface{}{},
	}
}

type Request struct {
	response   Response
	requestMap map[string]interface{}
}

func (req *Request) IsFailed() bool {
	return req.response.Failed()
}

func (req *Request) Fail() {
	req.response.SetStopped(true)
	req.response.SetFailed(true)
}

func (req *Request) Stop() {
	req.response.SetStopped(true)
}

func (req *Request) IsStopped() bool {
	return req.response.Stopped()
}
