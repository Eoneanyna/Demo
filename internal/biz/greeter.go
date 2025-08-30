package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/grpc"
)

type Greeter struct {
	Hello string
}
type GreeterStrem struct {
	grpc.ClientStream
}
type Service struct {
	Msg string
}

type GreeterRepo interface {
	CreateGreeter(context.Context, *Greeter) error
	UpdateGreeter(context.Context, *Greeter) (error, Service)
	UpdateStream(context.Context, *Greeter) (error, *GreeterStrem)
}

type GreeterUsecase struct {
	repo GreeterRepo
	log  *log.Helper
}

func NewGreeterUsecase(repo GreeterRepo, logger log.Logger) *GreeterUsecase {
	return &GreeterUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *GreeterUsecase) Create(ctx context.Context, g *Greeter) (error, Service) {
	return uc.repo.UpdateGreeter(ctx, g)
}

func (uc *GreeterUsecase) Update(ctx context.Context, g *Greeter) (error, *GreeterStrem) {

	err, strem := uc.repo.UpdateStream(ctx, g)

	if err != nil {
		return err, nil
	}

	return nil, strem
}
