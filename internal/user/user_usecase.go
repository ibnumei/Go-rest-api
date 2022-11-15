package user

import (
	"Go-rest-api/internal/domain"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserUsecase struct {
	db *gorm.DB
}

func NewUserUsecase(db *gorm.DB) *UserUsecase {
	return &UserUsecase{db: db}
}

func (uu UserUsecase) Register(c *gin.Context) {

	// bikin struct sendiri untuk keperluan input user, 
	// bisa berbeda dari struct  user (model struct) yang di sudah di define(yg sama seperti db)
	var userRequest UserRequest

	if err := c.ShouldBind(&userRequest);  err != nil {
		c.JSON(400, map[string]string{
			// "message": err.Error()  //for debug
			"message": "invalid input",
		})
		return
	}

	user, err  := domain.NewUser(userRequest.Name, userRequest.Email, userRequest.Password, userRequest.NoHp)
	if err != nil {
		c.JSON(400, map[string]string{
			"message": err.Error(),
		})
		return
	}
	if err := uu.db.Create(&user).Error; err != nil{
		c.JSON(500, map[string]string{
			"message": "error when create user",
		})
		return
	}

	token, err := user.GenerateJWT()
	if err != nil {
		c.JSON(500, map[string]string{
			"message": err.Error(),
			// "message": "error when generate token",
		})
		return
	}
	c.JSON(200, map[string]string{
		"token": token,
	})
 }

type UserRequest struct {
	Name     string
	Email    string
	Password string
	NoHp     string
}
