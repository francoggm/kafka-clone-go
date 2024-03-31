package broker

import (
	"net"
)

type consumer struct {
	conn  net.Conn
	topic *topic
}

func newConsumer(conn net.Conn, topic *topic) *consumer {
	return &consumer{
		conn:  conn,
		topic: topic,
	}
}


