package biz

import (
	"context"
	"demo/internal/data"
	"github.com/go-kratos/kratos/v2/log"
)

type NewsRepo interface {
	CreateNews(context.Context, *data.CreateNewsReq) (data.CreateNewsResp, error)
	GetNewsDetail(context.Context, *data.GetNewsDetailReq) (data.GetNewsDetailResp, error)
	GetNewsList(context.Context, *data.GetNewsListReq) (data.GetNewsListResp, error)
}

type NewsUsecase struct {
	repo NewsRepo
	rpc  *data.GRPCClient
	log  *log.Helper
}

func NewNewsUsecase(repo NewsRepo, rpc *data.GRPCClient, logger log.Logger) *NewsUsecase {
	return &NewsUsecase{repo: repo, rpc: rpc, log: log.NewHelper(logger)}
}

func (uc *NewsUsecase) CreateNews(ctx context.Context, req *data.CreateNewsReq) (data.CreateNewsResp, error) {
	//调用grpc服务
	resp, err := data.NewNewsRepo(log.GetLogger(), uc.rpc).CreateNews(ctx, &data.CreateNewsReq{
		Title:   req.Title,
		Content: req.Content,
	})
	if err != nil {
		return data.CreateNewsResp{}, err
	}
	return resp, nil
}

func (uc *NewsUsecase) GetNewsDetail(ctx context.Context, Id int32) (data.GetNewsDetailResp, error) {
	//调用grpc服务
	resp, err := data.NewNewsRepo(log.GetLogger(), uc.rpc).GetNewsDetail(ctx, &data.GetNewsDetailReq{
		Id: Id,
	})
	if err != nil {
		return data.GetNewsDetailResp{}, err
	}

	log.Info("GetNewsDetailResp:", resp)
	return data.GetNewsDetailResp{
		Id:         resp.Id,
		Title:      resp.Title,
		Content:    resp.Content,
		CreateTime: resp.CreateTime,
	}, nil
}

func (uc *NewsUsecase) GetNewsList(ctx context.Context, req *data.GetNewsListReq) (data.GetNewsListResp, error) {
	return uc.repo.GetNewsList(ctx, req)
}
