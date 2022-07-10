package mediator

import (
	"errors"
	"reflect"
)


var (
	registeredHandlers map[reflect.Type]any
)

func init() {
	registeredHandlers = make(map[reflect.Type]any)
}

func Register[Request,Result any](rType reflect.Type,handler RequestHandler[Request,Result]) error {
	registeredHandlers[rType] = handler
	return nil
}

func Send[Request,Result any](rType reflect.Type,r Request) (Result,error) {
	var a Result
	handler, ok := registeredHandlers[rType]
	if !ok {
		return a, errors.New("could not find a handler for this function")
	}
	switch handler := handler.(type) {
	case RequestHandler[Request,Result]:
		return handler.Handle(r)
	}
	return a, errors.New("Invalid handler")
}

type RequestHandler[Request,Result any] interface{
	Handle(request Request) (Result, error)
}