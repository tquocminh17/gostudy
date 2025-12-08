package main

import (
	"fmt"

	"github.com/tquocminh17/goapi/internal/user"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	userHandler := user.NewHandler()

	router.POST("/register", userHandler.Register)

	port := ":8080"
	fmt.Printf("Server started at %s\n", port)

	if err := router.Run(port); err != nil {
		fmt.Printf("Failed to start server: %v\n", err)
		panic(err)
	}
}
