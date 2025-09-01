package service

import (
	"context"
	v1 "demo/api/news"
	"demo/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/http"
)

type NewsService struct {
	uc  *biz.NewsUsecase
	log *log.Helper
}

func NewNewsService(uc *biz.NewsUsecase, logger log.Logger) *NewsService {
	return &NewsService{uc: uc, log: log.NewHelper(logger)}
}

// GetNewsById 根据ID获取新闻详情
func (s *NewsService) GetNewsById(ctx context.Context, req *v1.GetNewsByIdRequest) (*v1.GetNewsByIdResponse, error) {
	respStruct, err := s.uc.GetNewsDetail(ctx, req.Id)
	if err != nil {
		return &v1.GetNewsByIdResponse{}, err
	}

	return &v1.GetNewsByIdResponse{
		News: &v1.News{
			Id:         respStruct.Id,
			Title:      respStruct.Title,
			Content:    respStruct.Content,
			CreateTime: respStruct.CreateTime,
		}}, nil

}

// CreateNews 创建新闻
func (s *NewsService) CreateNews(ctx context.Context, req *v1.CreateNewsRequest) (*v1.CreateNewsResponse, error) {
	respStruct, err := s.uc.CreateNews(ctx, &biz.CreateNewsReq{
		Title:   req.Title,
		Content: req.Content,
	})
	if err != nil {
		return &v1.CreateNewsResponse{}, err
	}

	return &v1.CreateNewsResponse{
		Id: respStruct.Id,
	}, nil
}

// GetNewsList 创建新闻
func (s *NewsService) GetNewsList(ctx http.Context) error {
	//TODO
	//respStruct, err := s.uc.GetNewsList(ctx, &req)
	//if err != nil {
	//	return err
	//}
	//
	//ctx.JSON(200, respStruct)

	return nil
}
