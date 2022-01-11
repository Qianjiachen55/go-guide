package client

import (
	"gRPC_demo/gRPC_DEMO/global"
	"testing"
)

func TestGetOrder(t *testing.T) {
	global.ClientRun(GetAddOrder)
}



func TestSearchOrders(t *testing.T) {
	global.ClientRun(SearchOrders)
}

func TestUpdateOrder(t *testing.T) {
	global.ClientRun(UpdateOrder)
}