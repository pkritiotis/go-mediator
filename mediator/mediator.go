package mediator

import (
	"errors"
	"reflect"
	"sync"
)

var (
	registeredHandlers sync.Map
)

func init() {
	registeredHandlers = sync.Map{}
}

func Register[Request, Result any](rType reflect.Type, handler RequestHandler[Request, Result]) error {
	_, existed := registeredHandlers.LoadOrStore(rType, handler)
	if existed {
		return errors.New("the provided type is already registered to a handler")
	}
	return nil
}

func Send[Request, Result any](r Request) (Result, error) {
	var a Result
	handler, ok := registeredHandlers.Load(reflect.TypeOf(r))
	if !ok {
		return a, errors.New("could not find a handler for this function")
	}
	switch handler := handler.(type) {
	case RequestHandler[Request, Result]:
		return handler.Handle(r)
	}
	return a, errors.New("Invalid handler")
}

func SendCommand[Request any](rType reflect.Type, r Request) error {
	handler, ok := registeredHandlers.Load(rType)
	if !ok {
		return errors.New("could not find a handler for this function")
	}
	switch handler := handler.(type) {
	case CommandHandler[Request]:
		return handler.Handle(r)
	}
	return errors.New("Invalid handler")
}

type RequestHandler[Request, Result any] interface {
	Handle(request Request) (Result, error)
}

type CommandHandler[Request any] interface {
	Handle(request Request) error
}
