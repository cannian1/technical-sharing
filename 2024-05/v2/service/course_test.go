package service

import (
	"context"
	"testing"
)

type mockCourseClient struct {
}

func (c *mockCourseClient) LearnGo(ctx context.Context, req *Request) (map[string]any, error) {
	// http.Post 被打桩，不会真的发起请求

	return map[string]any{
		"msg": "Ok",
	}, nil
}

func (c *mockCourseClient) LearnJAVA(ctx context.Context, req *Request) (map[string]any, error) {
	return nil, nil
}

func (c *mockCourseClient) LearnC(ctx context.Context, req *Request) (map[string]any, error) {
	return nil, nil
}

func (c *mockCourseClient) LearnPython(ctx context.Context, req *Request) (map[string]any, error) {
	return nil, nil
}

func TestCourseServiceImpl_LearnGo(t *testing.T) {

	// 传入 mock 对象，测试的时候不会真的发起请求，会被打桩替换
	// 跑测试用例的时候接口保持幂等性
	mockService := NewCourseService(&mockCourseClient{})
	req := &Request{
		AAA: "aaa",
	}
	ctx := context.Background()

	resp, err := mockService.LearnGo(ctx, req)
	if resp == nil || resp.Msg != "Ok!" || err != nil {
		t.Errorf("TestCourseServiceImpl_LearnGo() failed")
	}

}
