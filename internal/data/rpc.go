package data

import (
	"context"
	"demo/internal/conf"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	mmd "github.com/go-kratos/kratos/v2/middleware/metadata"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	transgrpc "github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"google.golang.org/grpc"
	"time"
)

type GRPCClient struct {
	news *grpc.ClientConn
}

func NewGRPCClient(c *conf.Server, logger log.Logger) *GRPCClient {
	log := log.NewHelper(logger)
	sc := []constant.ServerConfig{
		*constant.NewServerConfig(c.Registry.Addr, c.Registry.Port),
	}

	cc := &constant.ClientConfig{
		NamespaceId:         c.Registry.Namespace, //namespace id
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "logs",
		CacheDir:            "nacos/cache",
		//RotateTime:          "1h",
		//MaxAge:              3,
		LogLevel: c.Registry.Loglevel,
	}

	// a more graceful way to create naming client
	cli, err := clients.NewNamingClient(
		vo.NacosClientParam{
			ClientConfig:  cc,
			ServerConfigs: sc,
		},
	)
	conn, err := NewRpcConn(cli, "demoserveice.NewsService", "news")
	if err != nil {
		log.Error(err.Error())
	}

	d := &GRPCClient{
		news: conn,
	}

	state := conn.GetState()
	log.Infof("连接状态: %v", state)

	return d
}
func NewRpcConn(cli naming_client.INamingClient, servicename string, group string) (*grpc.ClientConn, error) {
	conn, err := transgrpc.DialInsecure(
		context.Background(),
		//transgrpc.WithEndpoint("127.0.0.1:9000"),
		transgrpc.WithMiddleware(
			middleware.Chain(
				recovery.Recovery(),
				mmd.Client(),
			),
			tracing.Client(),
		),

		transgrpc.WithEndpoint("127.0.0.1:9001"),
		//transgrpc.WithDiscovery(registry.New(cli, registry.WithGroup(group))),
		transgrpc.WithTimeout(2*time.Second),
	)
	if err != nil {
		log.Error("grpc连接失败", err.Error())
		return nil, err
	}
	return conn, nil
}
