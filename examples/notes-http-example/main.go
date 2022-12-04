package main

import (
	"log"
	"net/http"

	"github.com/pkritiotis/go-mediator/examples/notes-http-example/infra"
)

func main() {
	infra.InitHttpServer()
	log.Fatal(http.ListenAndServe(":8080", nil))
}
