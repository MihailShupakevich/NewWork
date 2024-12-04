package main

import (
	"NewWork/internal/db"
	"github.com/gin-gonic/gin"
)

func main() {
	_, err := db.Connect()
	if err != nil {
		panic(err)
	}

	router := gin.Default()
	router.GET("/", func(c *gin.Context) {})
	router.Run(":8080")
}
