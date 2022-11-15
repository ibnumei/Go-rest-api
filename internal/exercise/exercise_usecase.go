package exercise

import (
	"Go-rest-api/internal/domain"
	"fmt"
	"strconv"
	"strings"

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

 func (eu ExerciseUsecase) GetScore(c *gin.Context) {
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
			"message": "not found",
		})
		return
	}
	userID := c.Request.Context().Value("user_id").(int)
	var answers []domain.Answer
	err = eu.db.Where("exercise_id = ? AND user_id = ?", id, userID).Find(&answers).Error
	if err != nil {
		c.JSON(404, map[string]interface{}{
			"message": "not answered yet",
		})
		return 
	}
	mapQA := make(map[int]domain.Answer)
	for _, answer := range answers {
		fmt.Println(answer)
		mapQA[answer.QuestionID] = answer
	}
	fmt.Println(mapQA)

	var score int
	for _, question := range exercise.Questions {
		if strings.EqualFold(question.CorrectAnswer, mapQA[question.ID].Answer) {
			score += question.Score
		}
	}
	c.JSON(200, map[string]int{
		"score": score,
	})
 }