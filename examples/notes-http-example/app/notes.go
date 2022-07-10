package app

import (
	"context"
	"github.com/google/uuid"
	"github.com/pkritiotis/go-mediate/mediator"
	"time"
)

func init() {
	mediator.Register(NewGetAllNotesRequestHandler())
}

// GetAllNotesRequest contains the request params
type GetAllNotesRequest struct {
	Ctx context.Context
}

// GetAllNotesResult is the result of the GetAllNotesRequest Query
type GetAllNotesResult struct {
	ID        uuid.UUID
	Name      string
	Contents  string
	CreatedAt time.Time
}

//GetAllNotesRequestHandler Contains the dependencies of the Handler
type GetAllNotesRequestHandler interface {
	Handle() ([]GetAllNotesResult, error)
}

type getAllNotesRequestHandler struct {
}

//NewGetAllNotesRequestHandler Handler constructor
func NewGetAllNotesRequestHandler() mediator.RequestHandler[GetAllNotesRequest, []GetAllNotesResult] {
	return getAllNotesRequestHandler{}
}

//Handle Handles the query
func (h getAllNotesRequestHandler) Handle(request GetAllNotesRequest) ([]GetAllNotesResult, error) {

	result := []GetAllNotesResult{
		{
			ID:        uuid.New(),
			Name:      "sample name",
			Contents:  "sample content",
			CreatedAt: time.Now(),
		},
	}

	return result, nil
}
