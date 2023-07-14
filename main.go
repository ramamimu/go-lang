package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func main() {
	fmt.Println("Hello World")
	router := gin.Default()

	user:= User{
		ID: 1,
		Name: "Rama",
	}

	fmt.Println(user)

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	})

	router.Run(":8080")
}
