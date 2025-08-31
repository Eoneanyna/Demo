package service

import (
	"demo/internal/biz"
	"encoding/json"
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
func (s *NewsService) GetNewsById(ctx http.Context) error {
	//TODO 这里先写死，后续改成传参
	respStruct, err := s.uc.GetNewsDetail(ctx, 1)
	if err != nil {
		return err
	}

	resp, _ := json.Marshal(respStruct)
	ctx.JSON(200, resp)

	return nil
}
