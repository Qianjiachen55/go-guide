package server

import (
	"context"
	"gRPC_demo/gRPC_DEMO/pb"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ProductServer struct {
	productMap map[string]*pb.Product
}

func (s *ProductServer) AddProduct(ctx context.Context, in *pb.Product) (*pb.ProductID, error) {
	out := uuid.New()
	in.Id = out.String()
	if s.productMap == nil {
		s.productMap = make(map[string]*pb.Product)
	}
	s.productMap[in.Id] = in
	return &pb.ProductID{Value: in.Id}, status.New(codes.OK, "").Err()
}

func (s *ProductServer) GetProduct(ctx context.Context, in *pb.ProductID) (*pb.Product,error){
	value,exists := s.productMap[in.Value]
	if exists{
		return value,status.New(codes.OK,"").Err()
	}

	return nil,status.Errorf(codes.NotFound,"Product does not exist",in.Value)
}
