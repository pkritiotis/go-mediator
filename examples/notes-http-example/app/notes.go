package app

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/pkritiotis/go-mediator/mediator"
)

func init() {
	err := mediator.Register[GetAllNotesRequest, []GetAllNotesResult](getAllNotesRequestHandler{})
	if err != nil {
		panic(err)
	}
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

// GetAllNotesRequestHandler Contains the dependencies of the Handler
type GetAllNotesRequestHandler interface {
	Handle() ([]GetAllNotesResult, error)
}

type getAllNotesRequestHandler struct {
}

// Handle Handles the query
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
