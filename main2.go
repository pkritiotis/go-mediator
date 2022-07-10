package main

import (
	"fmt"
	"github.com/pkritiotis/go-mediate/mediator"
	"reflect"
)

func main() {
	//Register the handler that will server `Request` requests
	mediator.Register(reflect.TypeOf(Request{}), NewSampleHandler())
	mediator.Send[Request, Result](Request{})
}

type Request struct {
}

type Result struct {
}

type SampleHandler struct {
}

func (s SampleHandler) Handle(request Request) (Result, error) {
	fmt.Println("Hello World")
	return Result{}, nil
}

func NewSampleHandler() mediator.RequestHandler[Request, Result] {
	return SampleHandler{}
}
