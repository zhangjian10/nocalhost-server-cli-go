package api

import (
	"os"

	"github.com/go-resty/resty/v2"
)

func getV1Request() *resty.Request {
	return resty.
		New().
		SetBaseURL(os.Getenv("NH_SERVER_HOST")).
		R().
		SetHeader("Content-Type", "application/json")
}
