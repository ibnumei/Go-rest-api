package usecase

import (
	// "Go-rest-api/internal/domain"

	// "github.com/gin-gonic/gin"
	"Go-rest-api/internal/domain"
	"context"
	"errors"
	// "gorm.io/gorm"
)

type UserGetter interface {
	GetByID(ctx context.Context, ID int) (domain.User  ,error)
}

type UserUsecase struct {
	// db *gorm.DB
	repo UserGetter
}

func NewUserUsecase(repo UserGetter) *UserUsecase {
	return &UserUsecase{repo: repo}
}

func (uu UserUsecase) GetUserByID(ctx context.Context, userID int) (domain.User, error) {
	// var user domain.User
	// err := uu.db.Where("id = ? ", userID).Take(&user).Error
	// return user, err

	if userID <= 0 {
		return domain.User{}, errors.New("Invalid User")
	}
	return  uu.repo.GetByID(ctx, userID)
}

// func (uu UserUsecase) Register(c *gin.Context) {

// 	// bikin struct sendiri untuk keperluan input user, 
// 	// bisa berbeda dari struct  user (model struct) yang di sudah di define(yg sama seperti db)
// 	var userRequest UserRequest

// 	if err := c.ShouldBind(&userRequest);  err != nil {
// 		c.JSON(400, map[string]string{
// 			// "message": err.Error()  //for debug
// 			"message": "invalid input",
// 		})
// 		return
// 	}

// 	user, err  := domain.NewUser(userRequest.Name, userRequest.Email, userRequest.Password, userRequest.NoHp)
// 	if err != nil {
// 		c.JSON(400, map[string]string{
// 			"message": err.Error(),
// 		})
// 		return
// 	}
// 	if err := uu.db.Create(&user).Error; err != nil{
// 		c.JSON(500, map[string]string{
// 			"message": "error when create user",
// 		})
// 		return
// 	}

// 	token, err := user.GenerateJWT()
// 	if err != nil {
// 		c.JSON(500, map[string]string{
// 			"message": err.Error(),
// 			// "message": "error when generate token",
// 		})
// 		return
// 	}
// 	c.JSON(200, map[string]string{
// 		"token": token,
// 	})
//  }

// type UserRequest struct {
// 	Name     string
// 	Email    string
// 	Password string
// 	NoHp     string
// }

