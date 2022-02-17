package api

import (
	"fmt"
	"os"

	"github.com/go-resty/resty/v2"
)

type BaseResult struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type Request struct {
	*resty.Request
}

const (
	V1 = "/v1"
	V2 = "v2"
)

func request(version string) *Request {

	r := &Request{
		resty.
			New().
			SetBaseURL(os.Getenv("NH_SERVER_HOST") + version).
			R(),
	}

	return r
}

var token string

func setToken(t string) {
	token = t
}

func (r *Request) Execute(method, url string) (*BaseResult, error) {

	var result *BaseResult
	var err error

	request := r.
		SetResult(&result)

	if token != "" {
		request = r.
			SetHeader("authorization", "Bearer "+token)
	}

	res, err := request.Execute(method, url)

	if !res.IsSuccess() {
		err = fmt.Errorf("%d %s", res.StatusCode(), res.Status())
	}
	return result, err
}
