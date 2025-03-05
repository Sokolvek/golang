package main

import (
	"playground/internal/routes"
	"playground/storage"
)

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @host      localhost:8080
// @BasePath  /
func main() {
	storage.InitDB()
	routes.Init()
}
