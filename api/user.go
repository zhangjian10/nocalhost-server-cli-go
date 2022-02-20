package api

import (
	"log"
	"nocalhost/server/utils/assert"

	"github.com/go-resty/resty/v2"
)

type loginResult struct {
	Token string `json:"token"`
}

func Login(email string, password string) (*loginResult, error) {
	request := &Request{
		request(V1).SetBody(map[string]interface{}{"email": email, "password": password}),
	}

	var r loginResult

	_, err := request.Execute(resty.MethodPost, "/login", &r)

	if err != nil {
		return nil, err
	}

	assert.NotEmpty(r, "login")

	log.Printf("\"%s\" login success", email)

	return &r, nil
}

var UserId int = 1
