package exercise

import (
	"Go-rest-api/internal/domain"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

 type ExerciseUsecase struct {
	db *gorm.DB
 }

 func NewExerciseUsecase(db *gorm.DB) *ExerciseUsecase {
	return &ExerciseUsecase{db: db}
 }

 func (eu ExerciseUsecase) GetExerciseByID(c *gin.Context) {
	idString := c.Param("id")
	id, err  := strconv.Atoi(idString)
	if err != nil {
		c.JSON(400, map[string]string{
			"message": "Invalid id",
		})
		return
	}
	var exercise domain.Exercise
	err = eu.db.Where("id = ?", id).Preload("Questions").Take(&exercise).Error
	if err != nil {
		c.JSON(400, map[string]string{
			"message": "Invalid id",
		})
		return
	}
	c.JSON(200, exercise)
 }