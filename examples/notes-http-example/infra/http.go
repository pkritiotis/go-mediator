package infra

import (
	"net/http"

	"github.com/pkritiotis/go-mediate/examples/notes-http-example/infra/notes"

	"github.com/gorilla/mux"
)

func InitHttpServer() {
	router := mux.NewRouter()

	http.Handle("/", router)
	router.HandleFunc("/notes", notes.GetAll).Methods("GET")
}
