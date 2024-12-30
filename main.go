// The main package implements a minimal example of the mediator
package main

import (
	"fmt"

	"github.com/pkritiotis/go-mediator/mediator"
)

func main() {
	//Register the handler that will server `TRequest` requests
	mediator.Register[Request, Result](SampleHandler{})
	mediator.Send[Request, Result](Request{value: "Hello World"})
}

// Request represents a typical request object
type Request struct {
	value string
}

// Result represents a typical result object
type Result struct {
}

// SampleHandler implements the Handler interface to serve Requests
type SampleHandler struct {
}

// Handle servers Request and returns a Result
func (s SampleHandler) Handle(request Request) (Result, error) {
	fmt.Println(request.value)
	return Result{}, nil
}
