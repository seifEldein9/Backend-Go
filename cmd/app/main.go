package main

import (
	"go-backend-app/internal/config"
	"go-backend-app/internal/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
     config.LoadEnv()

    r := gin.Default()

     r.GET("/users", handlers.GetAllUsers)
    r.POST("/users", handlers.CreateUser)
    r.POST("/login", handlers.LoginUser)

    port := ":8080"
    r.Run(port)
}
