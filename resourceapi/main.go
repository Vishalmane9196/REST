package main

import (
	"fmt"
	"ilmudata/resourceapi/controllers"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	//Load environment variables
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	//create gin instance
	r := gin.Default()
	todoRepo := controllers.New()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "welcome todo app",
		})
	})

	//group api
	r1 := r.Group("/api")
	{
		r1.POST("/todo", todoRepo.CreateTodo)
		r1.GET("/todo", todoRepo.GetTodos)
		r1.GET("/todo/:id", todoRepo.GetTodo)
		r1.PUT("/todo/:id", todoRepo.UpdateTodo)
		r1.DELETE("/todo/:id", todoRepo.DeleteTodo)
	}

	//start the web server
	r.Run("localhost:8080")
	fmt.Println("Server is running")
}
