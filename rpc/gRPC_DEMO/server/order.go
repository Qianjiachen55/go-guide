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
	"io"
	"log"
	"strings"
	"time"
)

type OrderServer struct {
	OrderMap map[string]*pb.Order
}

func (orderServer *OrderServer) AddOrder(ctx context.Context, order *pb.Order) (*wrappers.StringValue, error) {
	//order.Id = uuid.New().String()
	if orderServer.OrderMap == nil {
		orderServer.OrderMap = make(map[string]*pb.Order)
	}

	orderServer.OrderMap[order.Id] = order
	log.Println("add success", order.Id)
	return &wrappers.StringValue{Value: order.Id}, status.New(codes.OK, "").Err()
}

func (orderServer *OrderServer) GetOrder(ctx context.Context, orderId *wrappers.StringValue) (*pb.Order, error) {
	log.Println(cast.ToString(time.Now().Format(global.LAYOUT)))

	ord, exist := orderServer.OrderMap[orderId.Value]
	if exist {
		return ord, status.New(codes.OK, "").Err()
	}
	return nil, status.Errorf(codes.NotFound, "Order does not exist", orderId.Value)

}

func (orderServer *OrderServer) SearchOrders(searchQuery *wrappers.StringValue, stream pb.OrderManagement_SearchOrdersServer) error {
	global.PrintTime()
	fmt.Println("SearchOrders")
	for key,order := range orderServer.OrderMap{
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

func (orderServer *OrderServer) UpdateOrders(stream pb.OrderManagement_UpdateOrdersServer) error {

	ordersStr := "Updated Order IDs :"
	fmt.Println(ordersStr)
	for {
		order,err := stream.Recv()
		if err == io.EOF{
			return stream.SendAndClose(&wrappers.StringValue{Value: "Orders processed "+ ordersStr})
		}
		//order.Description = "after "+order.Id
		fmt.Println(order)
		fmt.Println("--------")
		orderServer.OrderMap[order.Id] = order
		log.Println("Order ID ",order.Id, ": Updated")
		ordersStr += order.Id + ","
	}
}

func (orderServer *OrderServer) ProcessOrders(stream pb.OrderManagement_ProcessOrdersServer)error{
	var combinedShipmentMap = make(map[string]*pb.CombinedShipment)
	for {
		val, err := stream.Recv()
		if err == io.EOF{
			//结束
			for _,comb := range combinedShipmentMap{
				stream.Send(comb)
			}
			return nil
		}
		if err != nil{
			//err
			log.Println(err)
			break
		}
		// process
		if val != nil{
			orderId := val.Value
			log.Printf("reading order: %+v\n", orderId)
			
			dest := orderServer.OrderMap[orderId].Destination
			
			shipment, exist := combinedShipmentMap[dest]
			if exist{
				ord := orderServer.OrderMap[orderId]
				shipment.OrdersList = append(shipment.OrdersList,ord)
				combinedShipmentMap[dest] = shipment
			}else {
				comShip := &pb.CombinedShipment{
					Id:         "cmb - "+(orderServer.OrderMap[orderId]).Destination,
					Status:     "Processed !",
					OrdersList: nil,
				}
				ord := orderServer.OrderMap[orderId]
				comShip.OrdersList = append(comShip.OrdersList, ord)
				combinedShipmentMap[dest] = comShip
				log.Println(len(comShip.OrdersList),comShip.GetId())
			}
			
		}
	}
	return nil
}