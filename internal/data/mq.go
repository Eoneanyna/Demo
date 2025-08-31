package data

import (
	"google.golang.org/grpc"
)

type MQClient struct {
	client *grpc.ClientConn
}

type MqData struct {
	Topic string
	Msg   string
	Key   string
}

//func (m *MQClient) SendMq(ctx context.Context, data *MqData) (res *pb.ProduceRepley, err error) {
//	mq := pb.NewProduceClient(m.client)
//	return mq.Produce(ctx, &pb.ProduceRequest{
//		Topic: data.Topic,
//		Msg:   data.Msg,
//		Key:   data.Key,
//	})
//}
