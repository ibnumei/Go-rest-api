package middleware

import (
	"Go-rest-api/internal/domain"
	"Go-rest-api/internal/user/usecase"
	"context"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
)

func WithAuth(userUsecase *usecase.UserUsecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(401, map[string]string{
				"message": "unauthorized",
			})
			c.Abort()
			return
		}

		if !strings.HasPrefix(authHeader, "Bearer ") {
			c.JSON(401, map[string]string{
				"message": "unauthorized",
			})
			c.Abort()
			return
		}

		auths := strings.Split(authHeader, " ")
		if len(auths) != 2 {
			c.JSON(401, map[string]string{
				"message": "unauthorized",
			})
			c.Abort()
			return
		}

		//decrypt jwt
		user := domain.User{}
		data, err := user.DecryptJWT(auths[1])
		fmt.Println(data)
		if err != nil {
			c.JSON(401, map[string]string{
				"message": "unauthorized",
			})
			c.Abort()
		}
		userID := int(data["user_id"].(float64))
		dbUser, err := userUsecase.GetUserByID(c.Request.Context(), userID)
		if err != nil || dbUser.ID == 0 {
			c.JSON(401, map[string]string{
				"message": "unauthorized",
			})
			c.Abort()
			return
		}
		ctxUserID := context.WithValue(c.Request.Context(), "user_id", int(data["user_id"].(float64)))
		c.Request = c.Request.WithContext(ctxUserID)
		
		c.Next()
	}
}