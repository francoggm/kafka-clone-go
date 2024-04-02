package model

import "net"

type BrokerConnection struct {
	Conn  net.Conn
	Mode  int
	Topic string
}
