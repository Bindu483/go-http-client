package auth

import (
	"encoding/json"
	"fmt"
	"github.com/Bindu483/go-http-client/utils"
	"log"
	"net/http"
	"os"
)

type Service struct {
	BaseUrl string
}

type Auth struct {
	JWT   string `json:"jwt"`
	Email string `json:"email"`
}

type LoginResponse struct {
	Meta MetaLoginResponse `json:"meta"`
}

type MetaLoginResponse struct {
	Token        string    `json:"token"`
	RefreshToken string    `json:"refreshToken"`
	User         *UserMeta `json:"user"`
}

type UserMeta struct {
	Email     string `json:"email"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	UserName  string `json:"username"`
}

func (s *Service) Login() (*Auth, error) {
	uname := os.Getenv("USERNAME")
	pwd := os.Getenv("PASSWORD")

	loginUrl := fmt.Sprintf("%s/api/v2/auth_server/login", s.BaseUrl)

	postBody, _ := json.Marshal(map[string]string{
		"username": uname,
		"password": pwd,
	})

	headers := map[string]string{
		"Content-Type": "application/json;charset=UTF-8",
		"Accept":       "application/json",
	}

	res, err := utils.MakeHTTPRequest(loginUrl, http.MethodPost, postBody, headers)
	if err != nil {
		return nil, err
	}

	var loginRes LoginResponse

	err = json.Unmarshal(res, &loginRes)
	if err != nil {
		log.Fatal(err)
	}

	return &Auth{
		JWT:   loginRes.Meta.Token,
		Email: loginRes.Meta.User.Email,
	}, nil
}
