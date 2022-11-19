package domain

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        int
	Name      string
	Email     string
	Password  string
	NoHp      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

var signature = []byte("mySignaturePrivateKey")

//validation request user from json
func NewUser(name, email, password, noHp string) (User, error) {
	if name == "" {
		return User{}, errors.New("Name can't be empty")
	}
	if email == "" {
		return User{}, errors.New("Email can't be empty")
	}
	if password == "" {
		return User{}, errors.New("Passworod can't be empty")
	}
	if noHp == "" {
		return User{}, errors.New("No Hp can't be empty")
	}

	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return User{
		Name:     name,
		Email:    email,
		Password: string(hash),
		NoHp:     noHp,
	}, nil
}

//generate jwt
func (u User) GenerateJWT() (string, error) {
	claims := jwt.MapClaims{
		"user_id": u.ID,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
		"iss":     "edspert",
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(signature)
	if err != nil {
		return "", err
	}
	return tokenString, err
}

  //decrypt jwt untuk auth akses data, middleware request
  func (u User) DecryptJWT(token string) (map[string]interface{}, error) {
	parsedToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("auth invalid")
		} 
		return signature, nil
	})

	data := make(map[string]interface{})
	if err != nil {
		return data, err
	}
	if !parsedToken.Valid {
		return data, errors.New("invalid token")
	}
	return parsedToken.Claims.(jwt.MapClaims), nil
}
