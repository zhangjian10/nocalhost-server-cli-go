package api

import (
	"os"

	"github.com/go-resty/resty/v2"
)

func getV1Request() *resty.Request {
	return getClient(v1).
		R().
		SetResult(&Result{})
}

func getV2Request() *resty.Request {
	return getClient(v2).
		R().
		SetResult(&Result{})
}

const (
	v1 = "v1"
	v2 = "v2"
)

var token string

func getClient(version string) *resty.Client {
	client := resty.
		New().
		SetBaseURL(os.Getenv("NH_SERVER_HOST") + version)

	if token != "" {
		client = client.
			SetHeader("authorization", "Bearer "+token)
	}

	return client
}

func setToken(newtoken string) {
	token = newtoken
}

type Result struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
