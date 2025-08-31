package data

import (
	"context"
	"demo/internal/conf"
	"github.com/go-kratos/kratos/v2/log"
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

	return d
}
func NewRpcConn(cli naming_client.INamingClient, servicename string, group string) (*grpc.ClientConn, error) {
	// 创建 gRPC 连接
	conn, err := transgrpc.DialInsecure(
		context.Background(),
		transgrpc.WithEndpoint("localhost:9001"), // 使用与 grpcurl 相同的地址
		transgrpc.WithTimeout(10*time.Second),    // 增加超时时间
	)

	if err != nil {
		log.Errorf("failed to dial grpc server: %v", err)
		return nil, err
	}
	return conn, nil
}
