package main

import (
	"log"
	"net/http"

	"github.com/tiltoin123/go-bookstore/pkg/routes"
)

func main() {

	routes.RegisterBookStoreRoutes()
	http.Handle("/", routes.Router)
	log.Fatal(http.ListenAndServe("localhost:9000", routes.Router))
}
