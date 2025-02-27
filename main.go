package main

import (
	"fmt"
	"net/http"
	"playground/routes"
	"playground/storage"
)

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @host      localhost:3000
// @BasePath  /
func main() {
	storage.InitDB()
	err := storage.Migrate()
	if err != nil {
		fmt.Printf("error caused by migrations %e", err)
	}
	routes.Init()
	http.ListenAndServe(":3000", routes.Mux)
}
