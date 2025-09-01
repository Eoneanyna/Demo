package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type NewsRepo interface {
	CreateNews(context.Context, *CreateNewsReq) (CreateNewsResp, error)
	GetNewsDetail(context.Context, *GetNewsDetailReq) (GetNewsDetailResp, error)
	GetNewsList(context.Context, *GetNewsListReq) (GetNewsListResp, error)
}

type NewsUsecase struct {
	repo NewsRepo
	log  *log.Helper
}

func NewNewsUsecase(repo NewsRepo, logger log.Logger) *NewsUsecase {
	return &NewsUsecase{repo: repo, log: log.NewHelper(logger)}
}

type CreateNewsReq struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}
type CreateNewsResp struct {
	Id int32 `json:"id"`
}

func (uc *NewsUsecase) CreateNews(ctx context.Context, req *CreateNewsReq) (CreateNewsResp, error) {
	//调用grpc服务
	resp, err := uc.repo.CreateNews(ctx, &CreateNewsReq{
		Title:   req.Title,
		Content: req.Content,
	})
	if err != nil {
		return CreateNewsResp{}, err
	}
	return resp, nil
}

type GetNewsDetailReq struct {
	Id int32 `json:"id"`
}

type GetNewsDetailResp struct {
	Id         int32  `json:"id"`
	Title      string `json:"title"`
	Content    string `json:"content"`
	CreateTime string `json:"create_time"`
}

func (uc *NewsUsecase) GetNewsDetail(ctx context.Context, Id int32) (GetNewsDetailResp, error) {
	//调用grpc服务
	resp, err := uc.repo.GetNewsDetail(ctx, &GetNewsDetailReq{
		Id: Id,
	})
	if err != nil {
		return GetNewsDetailResp{}, err
	}

	log.Info("GetNewsDetailResp:", resp)
	return resp, nil
}

type GetNewsListReq struct {
	PageNum int32 `json:"page_num"`
}

type GetNewsListResp struct {
	List []GetNewsDetailResp `json:"list"`
}

func (uc *NewsUsecase) GetNewsList(ctx context.Context, req *GetNewsListReq) (GetNewsListResp, error) {
	return uc.repo.GetNewsList(ctx, req)
}
