package client

import (
	"context"
	"demo/v3/service"
	"net/http"
)

// 服务的提供方不知道使用方会用到那些方法，不知道别人如何抽象
// 只需要把服务的实体类暴露出来，使用方自己去实现高级从抽象和屏蔽

func NewCourseClient() *CourseClient {
	return &CourseClient{}
}

type CourseClient struct {
}

func (c *CourseClient) LearnGo(ctx context.Context, req *service.Request) (map[string]any, error) {
	// ...
	_, err := http.Post("...", "application/json", nil)
	if err != nil {
		return nil, err
	}
	return map[string]any{
		"msg": "Ok",
	}, nil
}

func (c *CourseClient) LearnJAVA(ctx context.Context, req *service.Request) (map[string]any, error) {
	// ...
	_, err := http.Post("...", "application/json", nil)
	if err != nil {
		return nil, err
	}
	return map[string]any{
		"msg": "Ok",
	}, nil
}

func (c *CourseClient) LearnC(ctx context.Context, req *service.Request) (map[string]any, error) {
	// ...
	_, err := http.Post("...", "application/json", nil)
	if err != nil {
		return nil, err
	}
	return map[string]any{
		"msg": "Ok",
	}, nil
}

func (c *CourseClient) LearnPython(ctx context.Context, req *service.Request) (map[string]any, error) {
	// ...
	_, err := http.Post("...", "application/json", nil)
	if err != nil {
		return nil, err
	}
	return map[string]any{
		"msg": "Ok",
	}, nil
}
