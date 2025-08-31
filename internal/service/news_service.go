package service

import (
	"demo/internal/biz"
	"demo/internal/data"
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

// CreateNews 创建新闻
func (s *NewsService) CreateNews(ctx http.Context) error {
	//绑定参数
	req := data.CreateNewsReq{}
	if err := ctx.Bind(&req); err != nil {
	}
	respStruct, err := s.uc.CreateNews(ctx, &req)
	if err != nil {
		return err
	}

	ctx.JSON(200, respStruct)

	return nil
}

// GetNewsList 创建新闻
func (s *NewsService) GetNewsList(ctx http.Context) error {
	//绑定参数
	// 解析请求体中的JSON数据
	var req data.GetNewsListReq
	if err := ctx.Bind(&req); err != nil {
		log.Errorf("解析请求体失败: %v", err)
		return ctx.JSON(400, map[string]interface{}{
			"statusCode": 400,
			"message":    "请求参数错误: " + err.Error(),
			"data":       nil,
		})
	}

	respStruct, err := s.uc.GetNewsList(ctx, &req)
	if err != nil {
		return err
	}

	ctx.JSON(200, respStruct)

	return nil
}
