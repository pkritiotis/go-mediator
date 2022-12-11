package main

import (
	"log"
	"net/http"

	infra "github.com/pkritiotis/go-mediator/examples/notes-http-example/infra/http"
)

func main() {
	infra.InitServer()
	log.Fatal(http.ListenAndServe(":8080", nil))
}
