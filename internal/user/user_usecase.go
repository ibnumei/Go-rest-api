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

	if err := c.ShouldBind(&userRequest); err != nil {
		c.JSON(400, map[string]string{
			// "message": err.Error()  //for debug
			"message": "invalid input",
		})
		return
	}

	user, err := domain.NewUser(userRequest.Name, userRequest.Email, userRequest.Password, userRequest.NoHp)
	if err != nil {
		c.JSON(400, map[string]string{
			"message": err.Error(),
		})
		return
	}
	if err := uu.db.Create(&user).Error; err != nil {
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

type LoginParam struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (uu UserUsecase) Login(c *gin.Context) {
	var loginParam LoginParam

	if err := c.ShouldBind(&loginParam); err != nil {
		c.JSON(400, map[string]string{
			"message": "invalid input",
		})
		return
	}

	var user domain.User
	err := uu.db.Where("email = ?", loginParam.Email).Take(&user).Error

	if err != nil {
		c.JSON(400, map[string]string{
			"message": "wrong email",
		})
		return
	}

	if err := user.ComparePassword(loginParam.Password); err != nil {
		c.JSON(400, map[string]string{
			"message": "wrong password",
		})
		return
	}

	token, err := user.GenerateJWT()
	if err != nil {
		c.JSON(500, map[string]string{
			"message": err.Error(),
		})
		return
	}

	// return tipe-1
	// c.JSON(200, map[string]string{
	// 	"token": token,
	// })

	//return tipe-2
	//c.JSON(200, user)

	//return tipe-3
	c.JSON(200, map[string]string{
		"Name": user.Name,
		"NoHp": user.NoHp,
		"token": token,
	})
}
