package server

import (
	v1 "demo/api/news"
	"demo/internal/conf"
	"demo/internal/service"
	"demo/middleware/validate"
	"demo/pkg/encode"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/http"
)

// NewHTTPServer new a HTTP server.
func NewHTTPServer(c *conf.Server, newsService *service.NewsService) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			//tracing.Server(),
			logging.Server(log.GetLogger()),
			//metrics.Server(),
			validate.Validator(),
			tracing.Server(),
		),
	}

	//ctx := context.Background()
	//err := initTracer(ctx, "172.29.123.193:4317", "demo", "http")
	//if err != nil {
	//	fmt.Println("eee")
	//	log.Error("" + err.Error())
	//}

	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	opts = append(opts, http.ErrorEncoder(encode.ErrorEncoder))
	//opts = append(opts, http.ResponseEncoder(encode.ResponseEncoder))
	srv := http.NewServer(opts...)
	r := srv.Route("")
	r.GET("/checkHealth", func(ctx http.Context) error {
		ctx.JSON(200, map[string]string{"status": "UP"})
		return nil
	})
	//v1.RegisterGreeterHTTPServer(srv, GreeterService)
	//api.RegisterNewsHTTPServer(srv, newsService)
	v1.RegisterNewsServiceHTTPServer(srv, newsService)
	return srv
}
