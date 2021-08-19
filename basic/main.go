package main

import (
	"github.com/gin-gonic/gin"
)
func toKnow(c *gin.Context){
	c.JSON(200,gin.H{
		"message" : "trying to know about gin",
	})
}
func toWrite(c *gin.Context){
	name := c.Query("name")
	age := c.Query("age")
	c.JSON(200,gin.H{     // response object
		"name" : name,
		"age" : age,
	})

}
func main(){
	r:= gin.Default()
	r.GET("/know", toKnow)
	r.GET("/towrite", toWrite)
	r.Run(":8088")
}
