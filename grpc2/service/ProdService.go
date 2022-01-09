package service

import (
	"context"
)

type ProdService struct {

}

func (prodService *ProdService) GetProdStock (ctx context.Context, request *ProdRequest) (*ProdResponse, error){

	return &ProdResponse{
		ProdStock: 25,
	},nil
}

func (ProdService *ProdService)mustEmbedUnimplementedProdServiceServer(){

}