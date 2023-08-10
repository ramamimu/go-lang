package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}


type Mark struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Marks int `json:"marks"`
}

func main() {
	fmt.Println("Hello World")
	db, err := sql.Open("mysql", "root:root@/test")
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	result, err := db.Query("SELECT * FROM Marks")

	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		var user Mark
		err = result.Scan(&user.ID, &user.Name, &user.Marks)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println(user)
	}

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

	router.Run(":9090")
}
