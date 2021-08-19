package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func starting(c *gin.Context) {
	fmt.Println("start the coding")
	c.JSON(200,gin.H{
		"message" : "trying to know about gin",
	})
}

func main() {
	r := gin.Default()
	r.GET("/start", starting)
	r.Run(":8080")
}
