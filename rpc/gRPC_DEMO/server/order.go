package server

import (
	"context"
	"fmt"
	"gRPC_demo/gRPC_DEMO/global"
	"gRPC_demo/gRPC_DEMO/pb"
	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/spf13/cast"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"strings"
	"time"
)

type OrderServer struct {
	orderMap map[string]*pb.Order
}

func (orderServer *OrderServer) AddOrder(ctx context.Context, order *pb.Order) (*wrappers.StringValue, error) {
	//order.Id = uuid.New().String()
	if orderServer.orderMap == nil {
		orderServer.orderMap = make(map[string]*pb.Order)
	}

	orderServer.orderMap[order.Id] = order
	log.Println("add success", order.Id)
	return &wrappers.StringValue{Value: order.Id}, status.New(codes.OK, "").Err()
}

func (orderServer *OrderServer) GetOrder(ctx context.Context, orderId *wrappers.StringValue) (*pb.Order, error) {
	log.Println(cast.ToString(time.Now().Format(global.LAYOUT)))

	ord, exist := orderServer.orderMap[orderId.Value]
	if exist {
		return ord, status.New(codes.OK, "").Err()
	}
	return nil, status.Errorf(codes.NotFound, "Order does not exist", orderId.Value)

}

func (orderServer *OrderServer) SearchOrders(searchQuery *wrappers.StringValue, stream pb.OrderManagement_SearchOrdersServer) error {
	global.PrintTime()
	fmt.Println("SearchOrders")
	for key,order := range orderServer.orderMap{
		log.Print(key,order)
		for _, itemStr := range order.Items{
			log.Print(itemStr)
			if strings.Contains(itemStr, searchQuery.Value){
				err := stream.Send(order)
				if err != nil{
					return fmt.Errorf("error sending message to stream : %v",err)
				}
				log.Print("Matching Order Found : " + key)
				break
			}
		}
	}
	return nil
}
