package service

import (
	"demo/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/http"
	"strconv"
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
	// 从路径参数中获取 id
	idStr := ctx.Vars().Get("id")

	// 将字符串类型的 ID 转换为整数类型
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		// 如果转换失败，返回错误响应
		return ctx.JSON(400, map[string]interface{}{
			"statusCode": 400,
			"message":    "Invalid ID format",
			"data":       nil,
		})
	}

	respStruct, err := s.uc.GetNewsDetail(ctx, int32(id))
	if err != nil {
		return err
	}

	ctx.JSON(200, respStruct)

	return nil
}
