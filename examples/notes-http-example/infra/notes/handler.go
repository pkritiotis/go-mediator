package notes

import (
	"encoding/json"
	"fmt"
	"github.com/pkritiotis/go-mediate/examples/notes-http-example/app"
	"github.com/pkritiotis/go-mediate/mediator"
	"net/http"
)

//GetAll Returns all available notes
func GetAll(w http.ResponseWriter, r *http.Request) {
	req := app.GetAllNotesRequest{Ctx: r.Context()}
	crags, err := mediator.Send[app.GetAllNotesRequest, []app.GetAllNotesResult](req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, err.Error())
		return
	}

	err = json.NewEncoder(w).Encode(crags)
	if err != nil {
		return
	}
}
