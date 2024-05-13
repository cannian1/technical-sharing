package service

import (
	"context"
	"demo/v1/client"
)

type CourseService interface {
}

type CourseServiceImpl struct {
}

func (c *CourseServiceImpl) LearnGo(ctx context.Context, req *Request) (*Response, error) {
	// step1

	// step2

	// step3

	// Bad：每次调用这个函数都会创建一个新的 CourseClient 实例
	// 而且这个实例是不可复用的，因为它没有被保存下来
	// CourseClient 对象有自己的成员属性，它的生命周期应与 CourseService 保持一致
	cc := client.NewCourseClient()
	resp, err := cc.LearnGo(ctx, req)

	// step4

	// step5

	return &Response{
		Code: 200,
		Msg:  resp.Msg,
	}, err
}

type Response struct {
	Code int
	Msg  string
}

type Request struct {
	AAA string
}
