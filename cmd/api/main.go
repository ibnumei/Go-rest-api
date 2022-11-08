package main

import (
	"Go-rest-api/internal/database"
	"Go-rest-api/internal/exercise"

	"github.com/gin-gonic/gin"
)

func main() {
	route := gin.Default()
	
	route.GET("/hello", func(c *gin.Context)  {
		c.JSON(200, map[string]string{
			"message": "hellos",
		})
	})

	dbConn := database.NewDatabaseConn()
	eu := exercise.NewExerciseUsecase(dbConn)
	route.GET("/exercises/:id", eu.GetExerciseByID)

	route.Run(":1234")
}