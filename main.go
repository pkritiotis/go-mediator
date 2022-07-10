package main

import (
	"fmt"
	"github.com/pkritiotis/go-mediate/mediator"
	"reflect"
)


func main() {


	handler := NewSampleHandler()
	mediator.Register(reflect.TypeOf(Request{}),handler)
	mediator.Send[Request,Result](reflect.TypeOf(Request{}),Request{})

}

type Request struct{

}

type Result struct{

}

type SampleHandler struct{

}


func (s SampleHandler) Handle(request Request) (Result, error){
	fmt.Println("Hello World")
	return Result{},nil
}

//func NewSampleHandler() mediator.CommandHandler[string]{
//	return SampleHandler{}
//}

func NewSampleHandler() mediator.RequestHandler[Request, Result]{
	return SampleHandler{}
}
