package notes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pkritiotis/go-mediator/examples/notes-http-example/app"
	"github.com/pkritiotis/go-mediator/mediator"
)

// GetAll Returns all available notes
func GetAll(w http.ResponseWriter, r *http.Request) {
	req := app.GetAllNotesRequest{Ctx: r.Context()}
	notes, err := mediator.Send[app.GetAllNotesRequest, []app.GetAllNotesResult](req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, err.Error())
		return
	}

	err = json.NewEncoder(w).Encode(notes)
	if err != nil {
		return
	}
}
