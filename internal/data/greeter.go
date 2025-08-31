package data

//import (
//	"context"
//	"demo/internal/biz"
//	"fmt"
//	"github.com/go-kratos/kratos/v2/log"
//	v1 "gitlab.cqrb.cn/shangyou_mic/testpg/api/demoserveice/v1"
//	"google.golang.org/grpc"
//)
//
//type greeterRepo struct {
//	data      *Data
//	rpc       *GRPCClient
//	log       *log.Helper
//	baseinfos *BaseInfos
//}
//type Result struct {
//	Userid   int
//	Username string
//}
//
//var result Result
//
//// NewGreeterRepo .
//func NewGreeterRepo(data *Data, logger log.Logger, rpc *GRPCClient) biz.GreeterRepo {
//	return &greeterRepo{
//		data: data,
//		rpc:  rpc,
//		log:  log.NewHelper(logger),
//	}
//}
//
//func (r *greeterRepo) CreateGreeter(ctx context.Context, g *biz.Greeter) error {
//
//	return nil
//}
//
//func (r *greeterRepo) UpdateGreeter(ctx context.Context, g *biz.Greeter) (error, biz.Service) {
//	res, err := v1.NewGreeterClient(r.rpc.news).SayHello(ctx, &v1.HelloRequest{Name: "222"})
//
//	if err != nil {
//		fmt.Println("错误:", err.Error())
//		return err, biz.Service{}
//	}
//	msg := biz.Service{Msg: res.Message}
//
//	return nil, msg
//}
//func (r *greeterRepo) UpdateStream(ctx context.Context, g *biz.Greeter) (error, *biz.GreeterStrem) {
//	str, err := v1.NewDataStreamClient(r.rpc.news).Subscribe(ctx, &v1.SubscribeRequest{Topic: "222"})
//	//
//	GreeterStrem := new(biz.GreeterStrem)
//
//	if err != nil {
//		return err, GreeterStrem
//
//	}
//	b := str.(grpc.ClientStream)
//
//	GreeterStrem.ClientStream = b
//
//	//for {
//	//	res, err := str.Recv()
//	//	if err == io.EOF {
//	//		break
//	//	}
//	//	fmt.Println(res.Data)
//	//}
//	fmt.Println("end")
//	return nil, GreeterStrem
//}
