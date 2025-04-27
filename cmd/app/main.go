package main

import (
    "go-backend-app/internal/config"
    "go-backend-app/internal/handlers"

    "github.com/gin-gonic/gin"
)

func main() {
    // Load config
    config.LoadEnv()

    r := gin.Default()

    // Routes
    r.GET("/users", handlers.GetAllUsers)
    r.POST("/users", handlers.CreateUser)

    port := ":8080"
    r.Run(port)
}
