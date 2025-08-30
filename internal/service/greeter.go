package service

import (
	"context"
	"demo/internal/biz"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	v1 "gitlab.cqrb.cn/shangyou_mic/testpg/api/helloworld/v1"
)

// GreeterService is a greeter service.
type GreeterService struct {
	v1.UnimplementedGreeterServer

	uc  *biz.GreeterUsecase
	log *log.Helper
}
type SseService struct {
	uc  *biz.GreeterUsecase
	log *log.Helper
}

// NewGreeterService new a greeter service.
func NewGreeterService(uc *biz.GreeterUsecase, logger log.Logger) *GreeterService {
	return &GreeterService{uc: uc, log: log.NewHelper(logger)}
}
func NewSseService(uc *biz.GreeterUsecase, logger log.Logger) *SseService {
	return &SseService{uc: uc, log: log.NewHelper(logger)}
}

// SayHello implements helloworld.GreeterServer
func (s *GreeterService) SayHello(ctx context.Context, in *v1.HelloRequest) (*v1.HelloReply, error) {

	e, r := s.uc.Create(ctx, &biz.Greeter{Hello: "2222"})
	if e != nil {
		err2 := errors.New(500, e.Error(), "请求失败")
		return nil, err2
	}
	//for {
	//	res, err := str.Recv()
	//	if err == io.EOF {
	//		break
	//	}
	//	fmt.Println(res.Data)
	//}
	//if in.GetName() == "error" {
	//	return nil, v1.ErrorUserNotFound("user not found: %s", in.GetName())
	return &v1.HelloReply{Message: "Hello " + r.Msg}, nil
}
func (s *SseService) Sse(ctx context.Context, in *v1.SseRequest) (*v1.SseReply, error) {
	err, str := s.uc.Update(ctx, &biz.Greeter{Hello: "2222"})
	if err != nil {
		err2 := errors.New(500, err.Error(), "请求失败")
		return nil, err2
	}
	return &v1.SseReply{str}, nil
}
