package data

import (
	"context"
	"demo/internal/biz"
	v1 "demoserveice/api/news/v1"
	"github.com/go-kratos/kratos/v2/log"
	"time"
)

type newsRepo struct {
	log *log.Helper
	rpc *GRPCClient
}

func NewNewsRepo(logger log.Logger, rpc *GRPCClient) *newsRepo {
	return &newsRepo{log.NewHelper(logger), rpc}
}

func (r *newsRepo) CreateNews(ctx context.Context, news *biz.CreateNewsReq) (biz.CreateNewsResp, error) {
	res, err := v1.NewNewsServiceClient(r.rpc.news).CreateNews(ctx, &v1.CreateNewsRequest{Title: news.Title, Content: news.Content})
	if err != nil {
		r.log.Errorf("调用CreateNews接口失败: %v", err)
		return biz.CreateNewsResp{}, err
	}

	return biz.CreateNewsResp{
		Id: res.Id,
	}, nil
}

func (r *newsRepo) GetNewsDetail(ctx context.Context, req *biz.GetNewsDetailReq) (biz.GetNewsDetailResp, error) {
	res, err := v1.NewNewsServiceClient(r.rpc.news).GetNewsById(ctx, &v1.GetNewsByIdRequest{Id: req.Id})
	if err != nil {
		r.log.Errorf("调用GetNewsById接口失败: %v", err)
		return biz.GetNewsDetailResp{}, err
	}

	timeTemplate1 := "2006-01-02T15:04:05Z" //常规类型

	return biz.GetNewsDetailResp{
		Id:         res.Id,
		Title:      res.Title,
		Content:    res.Content,
		CreateTime: time.Unix(res.CreateTime, 0).Format(timeTemplate1),
	}, nil
}

func (r *newsRepo) GetNewsList(ctx context.Context, req *biz.GetNewsListReq) (biz.GetNewsListResp, error) {
	return biz.GetNewsListResp{}, nil
}
