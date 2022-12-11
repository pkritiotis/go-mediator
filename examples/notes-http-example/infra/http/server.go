package infra

import (
	"net/http"

	"github.com/pkritiotis/go-mediator/examples/notes-http-example/infra/http/notes"

	"github.com/gorilla/mux"
)

func InitServer() {
	router := mux.NewRouter()

	http.Handle("/", router)
	router.HandleFunc("/notes", notes.GetAll).Methods("GET")
}
