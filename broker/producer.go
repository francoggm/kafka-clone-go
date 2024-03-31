package broker

import (
	"net"
)

type producer struct {
	conn  net.Conn
	topic *topic
}

func newProducer(conn net.Conn, topic *topic) *producer {
	return &producer{
		conn:  conn,
		topic: topic,
	}
}

func (p *producer) produceMessage(key string, value string) {
	msg := newMessage(key, value)
	p.topic.addMessage(msg)
}
