package routes

import (
	"net/http"
	_ "playground/docs"
	"playground/handlers"

	httpSwagger "github.com/swaggo/http-swagger"
)

var Mux = http.NewServeMux()

func Init() {

	Mux.HandleFunc("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("/swagger/doc.json"),
	))
	Mux.HandleFunc("/users", handlers.FindAllUsers)
}
