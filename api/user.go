package api

import (
	"nocalhost/server/utils/assert"

	"github.com/go-resty/resty/v2"
)

type loginResult struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}

func Login(email string, password string) {
	request := &Request{
		request(V1).SetBody(map[string]interface{}{"email": email, "password": password}),
	}

	var r loginResult

	res := request.Execute(resty.MethodPost, "/login", &r)

	if res.isSuccess() {

		assert.NotEmpty(r, "login")

		SetToken(r.Token)
	}
}

var UserId int = 1
