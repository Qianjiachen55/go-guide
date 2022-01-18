package config

const (
	RMQADDR = "amqp://guest:guest@localhost:5672/"
	EXCHANGENAME = "direct_exchange"
	CONSUMERCNT = 3
)

var RoutingKeys [4]string = [4]string{"info", "debug", "warn", "error"}

