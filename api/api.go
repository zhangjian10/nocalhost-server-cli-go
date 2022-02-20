package api

import (
	"fmt"

	"github.com/go-resty/resty/v2"
	"github.com/mitchellh/mapstructure"
)

type baseResult struct {
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
			SetBaseURL(hostname + version).
			R(),
	}

	return r
}

var token, hostname string

func SetParameters(h, t string) {
	hostname = h
	token = t
}
func (r *baseResult) isSuccess() bool {
	return r.Code == 0
}

func (r *baseResult) error() error {
	return fmt.Errorf("service error, code:%d, message:%s", r.Code, r.Message)
}

func (r *Request) Execute(method, url string, output interface{}) (*baseResult, error) {

	var result *baseResult
	var err error

	request := r.
		SetResult(&result)

	if token != "" {
		request = r.
			SetHeader("authorization", "Bearer "+token)
	}

	res, err := request.Execute(method, url)

	if err != nil {
		return nil, err
	}

	if !res.IsSuccess() {
		return nil, fmt.Errorf("request fail, code:%v, status:%v, body:%v", res.StatusCode(), res.Status(), string(res.Body()))
	}

	if !result.isSuccess() {
		return nil, result.error()
	}

	if output != nil {
		config := &mapstructure.DecoderConfig{TagName: "json", Result: &output}
		decoder, _ := mapstructure.NewDecoder(config)

		decoder.Decode(result.Data)
	}

	return result, nil
}
