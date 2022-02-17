package api

import (
	"fmt"

	"github.com/go-resty/resty/v2"
)

func Login(email string, password string) {
	request := &Request{
		request(V1).SetBody(map[string]interface{}{"email": email, "password": password}),
	}

	resp, err := request.Execute(resty.MethodPost, "/login")

	fmt.Println(resp.Data, err)
}
