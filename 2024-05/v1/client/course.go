package client

import (
	"context"
	"demo/v1/service"
	"net/http"
)

type CourseClient interface {
	LearnGo(ctx context.Context, req *service.Request) (*Response, error)
	LearnJAVA()
	LearnC()
	LearnPython()
	// ...
}

type courseClientImpl struct {
}

func NewCourseClient() CourseClient {
	return &courseClientImpl{}
}

func (c *courseClientImpl) LearnGo(ctx context.Context, req *service.Request) (*Response, error) {
	// ...
	_, err := http.Post("...", "application/json", nil)
	if err != nil {
		return nil, err
	}
	return &Response{
		Msg: "Ok",
	}, nil
}

func (c *courseClientImpl) LearnJAVA() {

}

func (c *courseClientImpl) LearnC() {

}

func (c *courseClientImpl) LearnPython() {

}

type Request struct {
	AAA string
	BBB string
}

type Response struct {
	Msg string
}
