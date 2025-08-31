package api

import (
	"demo/internal/service"
	"github.com/go-kratos/kratos/v2/transport/http"
)

func RegisterNewsHTTPServer(s *http.Server, newsService *service.NewsService) {
	r := s.Route("/demo")
	r.GET("/news/{id}", newsService.GetNewsById)
	//r.POST("/news", newsService.CreateNews)
}
