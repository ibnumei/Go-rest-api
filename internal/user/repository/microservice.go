package repository

import (
	"Go-rest-api/internal/domain"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)
const (
	getUsersURL= "/internal/users"
)
type McsUserRepo struct {
	Hostname string
	Username string
	Password string
	Client *http.Client
}

func NewMcsUserRepo() *McsUserRepo {
	client := http.Client{
		Timeout: 30 * time.Second,
	}
	return &McsUserRepo{
		Hostname : "http://localhost:8080",
		Username : "user",
		Password: "abcd1234",
		Client: &client,
	}
}

func (mcsUsecase McsUserRepo) GetByID(ctx context.Context, ID int) (domain.User, error) {
	// http://localhost:8080/internal/users/:id
	url := fmt.Sprintf("%s%s/%d", mcsUsecase.Hostname, getUsersURL, ID)
	//NewRequestWithContext context bisa sampe ke ms lain, bisa dibuat tracing. 
	//NewRequest tidak bisa sampe ke ms lain
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)

	var user domain.User
	if err != nil {
		return user, err
	}

	req.SetBasicAuth(mcsUsecase.Username, mcsUsecase.Password)

	//hit endpoint
	log.Println("hit to", req.Method, req.URL)
	resp, err :=  mcsUsecase.Client.Do(req)
	if err != nil {
		log.Println(err)
		return user, err
	}
	if resp.StatusCode != http.StatusOK {
		log.Println(err)
		return user, err
	}
	err = json.NewDecoder(resp.Body).Decode(&user)
	return user, err
}
