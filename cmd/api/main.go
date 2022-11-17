package main

import (
	"Go-rest-api/internal/database"
	"Go-rest-api/internal/exercise"
	"Go-rest-api/internal/middleware"

	// "Go-rest-api/internal/user"
	"Go-rest-api/internal/user/repository"
	"Go-rest-api/internal/user/usecase"

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
	userRepo := repository.NewUserDBRepo(dbConn)
	uu := usecase.NewUserUsecase(userRepo)

	//exercise endpoint
	route.GET("/exercises/:id", middleware.WithAuth(uu), eu.GetExerciseByID)

	route.GET("/exercises/:id/scores", middleware.WithAuth(uu), eu.GetScore)

	//user endpoint
	// route.POST("/register", uu.Register)
	route.Run(":1234")
}