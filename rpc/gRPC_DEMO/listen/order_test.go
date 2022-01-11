package listen

import (
	"gRPC_demo/gRPC_DEMO/global"
	"testing"
)

func TestOrderListen(t *testing.T) {
	//go OrderListen()
	//
	//select {
	//
	//}

	global.ServerRun(OrderListen)
}
