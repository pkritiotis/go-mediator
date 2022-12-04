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

type key[TRequest any, TResult any] struct {
	reqType reflect.Type
	resType reflect.Type
}

// Register registers the provided request handler to be used for the corresponding requests
func Register[TRequest any, TResult any](handler RequestHandler[TRequest, TResult]) error {
	var req TRequest
	var res TResult
	k := key[TRequest, TResult]{
		reqType: reflect.TypeOf(req),
		resType: reflect.TypeOf(res),
	}

	_, existed := registeredHandlers.LoadOrStore(k, handler)
	if existed {
		return errors.New("the provided type is already registered to a handler")
	}
	return nil
}

// SendRequest processes the provided request and returns the produced result
func SendRequest[TRequest any, TResult any](r TRequest) (TResult, error) {
	var zeroRes TResult
	var k key[TRequest, TResult]
	handler, ok := registeredHandlers.Load(reflect.TypeOf(k))
	if !ok {
		return zeroRes, errors.New("could not find zeroRes handler for this function")
	}
	switch handler := handler.(type) {
	case RequestHandler[TRequest, TResult]:
		return handler.Handle(r)
	}
	return zeroRes, errors.New("Invalid handler")
}

// RequestHandler handles TRequest and returns TResult
type RequestHandler[TRequest any, TResult any] interface {
	Handle(request TRequest) (TResult, error)
}

// SendCommand processes the provided Request
func SendCommand[TRequest any](r TRequest) error {
	var req TRequest
	handler, ok := registeredHandlers.Load(reflect.TypeOf(req))
	if !ok {
		return errors.New("could not find zeroRes handler for this function")
	}
	switch handler := handler.(type) {
	case CommandHandler[TRequest]:
		return handler.Handle(r)
	}
	return errors.New("Invalid handler")
}

// CommandHandler handles TRequest
type CommandHandler[TRequest any] interface {
	Handle(request TRequest) error
}
