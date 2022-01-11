package service

import (
	pb "productInfo/service/ecommerce"
)

type server struct {
	productMap map[string]*pb.Product
}
