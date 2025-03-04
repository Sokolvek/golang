package routes

import (
	_ "playground/docs"
	"playground/internal/handlers"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Init() {
	r := gin.Default()
	r.POST("/users", handlers.CreateUser)
	r.GET("/users", handlers.FindAllUsers)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.Run()
}
