package routes

import (
	"net/http"

	"github.com/tiltoin123/go-bookstore/pkg/controllers"
)

var Router = http.NewServeMux()

func RegisterBookStoreRoutes() {
	Router.HandleFunc("POST /book/", controllers.CreateBook)
	Router.HandleFunc("PUT /book/{id}", controllers.UpdateBook)
	Router.HandleFunc("GET /book/", controllers.GetBooks)
	Router.HandleFunc("GET /book/{id}", controllers.GetBookById)
	Router.HandleFunc("DELETE /book/{id}", controllers.DeleteBook)
}
