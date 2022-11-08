package main

import "github.com/gin-gonic/gin"

func main() {
	route := gin.Default()
	
	route.GET("/hello", func(c *gin.Context)  {
		c.JSON(200, map[string]string{
			"message": "hellos",
		})
	})
	route.Run(":1234")
}