package main

import (
	"Go-rest-api/internal/database"
	"Go-rest-api/internal/exercise"
	"Go-rest-api/internal/middleware"
	"Go-rest-api/internal/user"

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
	uu := user.NewUserUsecase(dbConn)

	//exercise endpoint
	route.GET("/exercises/:id", middleware.WithAuth(), eu.GetExerciseByID)

	route.GET("/exercises/:id/scores", middleware.WithAuth(), eu.GetScore)

	//user endpoint
	route.POST("/register", uu.Register)
	route.Run(":1234")
}