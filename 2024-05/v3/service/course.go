package service

import (
	"context"
)

// 使用方只要用到 LearnGo 方法，就近定义 courseProxy 接口，屏蔽其他底层方法
type courseProxy interface {
	LearnGo(ctx context.Context, req *Request) (*Response, error)
	// 需要扩展其他方法时，只需要在这里和 mock 定义，仍然不会暴露底层方法，聚焦在业务逻辑上
	//LearnPiano(ctx context.Context, req *Request) (*Response, error)
}

type courseServiceImpl struct {
	courseClient courseProxy
	// 其他的需要长生命周期的成员属性，比如数据库连接等也可以在这里定义，由下面的构造器初始化
}

// 除了 mock 之外，还可以在这里定义其他的 client 实现，比如 grpcClient、httpClient 等
// 或者不同的 client 实现，比如 courseClientV1、courseClientV2，又或者在不同的地域下使用不同的 client 实现
func NewCourseService(courseClient courseProxy) courseProxy {
	return &courseServiceImpl{
		courseClient: courseClient,
	}
}

func (c *courseServiceImpl) LearnGo(ctx context.Context, req *Request) (*Response, error) {
	// step1

	// step2

	// step3

	// 在调用的时候不会出现模块外的方法干扰
	resp, err := c.LearnGo(ctx, req)

	// step4

	// step5

	return &Response{
		Code: 200,
		Msg:  resp.Msg,
	}, err
}

//func (c *courseServiceImpl) LearnPiano(ctx context.Context, req *Request) (*Response, error) {
//	return &Response{
//		Code: 200,
//		Msg:  "不依附于 client 层的方法",
//	}, nil
//}

type Response struct {
	Code int
	Msg  string
}

type Request struct {
	AAA string
}
