package service

import (
	"context"
	"demo/v2/client"
)

type CourseService interface {
	LearnGo(ctx context.Context, req *Request) (*Response, error)
}

type CourseServiceImpl struct {
	courseClient client.CourseClient
	// 其他的需要长生命周期的成员属性，比如数据库连接等也可以在这里定义，由下面的构造器初始化
}

func NewCourseService(courseClient client.CourseClient) CourseService {
	return &CourseServiceImpl{
		courseClient: courseClient,
	}
}

func (c *CourseServiceImpl) LearnGo(ctx context.Context, req *Request) (*Response, error) {
	// step1

	// step2

	// step3

	// Good：CourseClient 对象是可复用的
	// 使用 mock 替换真实的 CourseClient 对象，可以在测试的时候规避不幂等的接口调用
	// 从而关注其他 step 的业务逻辑
	resp, err := c.courseClient.LearnGo(ctx, req)

	// step4

	// step5

	return &Response{
		Code: 200,
		Msg:  resp["Msg"].(string),
	}, err
}

type Response struct {
	Code int
	Msg  string
}

type Request struct {
	AAA string
}
