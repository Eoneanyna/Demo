package data

import (
	"context"
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

type CreateNewsReq struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}
type CreateNewsResp struct {
	Id int32 `json:"id"`
}

func (r *newsRepo) CreateNews(ctx context.Context, news *CreateNewsReq) (CreateNewsResp, error) {
	return CreateNewsResp{}, nil
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

func (r *newsRepo) GetNewsDetail(ctx context.Context, req *GetNewsDetailReq) (GetNewsDetailResp, error) {
	res, err := v1.NewNewsServiceClient(r.rpc.news).GetNewsById(ctx, &v1.GetNewsByIdRequest{Id: req.Id})
	if err != nil {
		r.log.Errorf("调用GetNewsById接口失败: %v", err)
		return GetNewsDetailResp{}, err
	}

	timeTemplate1 := "2006-01-02T15:04:05Z" //常规类型

	return GetNewsDetailResp{
		Id:         res.News.Id,
		Title:      res.News.Title,
		Content:    res.News.Content,
		CreateTime: time.Unix(res.News.CreateTime, 0).Format(timeTemplate1),
	}, nil
}

type GetNewsListReq struct {
	PageNum int32 `json:"page_num"`
}

type GetNewsListResp struct {
	List []GetNewsDetailResp `json:"list"`
}

func (r *newsRepo) GetNewsList(ctx context.Context, req *GetNewsListReq) (GetNewsListResp, error) {
	return GetNewsListResp{}, nil
}
