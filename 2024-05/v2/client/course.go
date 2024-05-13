package client

import (
	"context"
	"demo/v2/service"
	"net/http"
)

type CourseClient interface {
	LearnGo(ctx context.Context, req *service.Request) (map[string]any, error)
	LearnJAVA(ctx context.Context, req *service.Request) (map[string]any, error)
	LearnC(ctx context.Context, req *service.Request) (map[string]any, error)
	LearnPython(ctx context.Context, req *service.Request) (map[string]any, error)
	// ...
}

type courseClientImpl struct {
}

func NewCourseClient() CourseClient {
	return &courseClientImpl{}
}

func (c *courseClientImpl) LearnGo(ctx context.Context, req *service.Request) (map[string]any, error) {
	// ...
	_, err := http.Post("...", "application/json", nil)
	if err != nil {
		return nil, err
	}
	return map[string]any{
		"msg": "Ok",
	}, nil
}

func (c *courseClientImpl) LearnJAVA(ctx context.Context, req *service.Request) (map[string]any, error) {
	// ...
	_, err := http.Post("...", "application/json", nil)
	if err != nil {
		return nil, err
	}
	return map[string]any{
		"msg": "Ok",
	}, nil
}

func (c *courseClientImpl) LearnC(ctx context.Context, req *service.Request) (map[string]any, error) {
	// ...
	_, err := http.Post("...", "application/json", nil)
	if err != nil {
		return nil, err
	}
	return map[string]any{
		"msg": "Ok",
	}, nil
}

func (c *courseClientImpl) LearnPython(ctx context.Context, req *service.Request) (map[string]any, error) {
	// ...
	_, err := http.Post("...", "application/json", nil)
	if err != nil {
		return nil, err
	}
	return map[string]any{
		"msg": "Ok",
	}, nil
}

type Request struct {
	AAA string
	BBB string
}
