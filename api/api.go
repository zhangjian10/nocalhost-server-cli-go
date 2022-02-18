package api

import (
	"fmt"
	"os"

	"github.com/go-resty/resty/v2"
	"github.com/mitchellh/mapstructure"
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
	V2 = "/v2"
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

func SetToken(t string) {
	token = t
}

func (r BaseResult) isSuccess() bool {
	return r.Code == 0
}

func (r BaseResult) error() string {
	return fmt.Sprintf("[ServiceError]\n code:%v\n message:%v", r.Code, r.Message)
}

func (r *Request) Execute(method, url string, output interface{}) *BaseResult {

	var result *BaseResult
	var err error

	request := r.
		SetResult(&result)

	if token != "" {
		request = r.
			SetHeader("authorization", "Bearer "+token)
	}

	res, err := request.Execute(method, url)

	if err != nil {
		panic(err)
	}

	if !res.IsSuccess() {
		panic(fmt.Errorf("%d %s", res.StatusCode(), res.Status()))
	}

	if !result.isSuccess() {
		panic(result.error())
	}

	if output != nil {
		mapstructure.Decode(result.Data, output)
	}

	return result
}
