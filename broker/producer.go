package broker

import (
	"bufio"
	"net"
)

type producer struct {
	connected bool
	conn      net.Conn
	topic     *topic
}

func newProducer(conn net.Conn, topic *topic) *producer {
	return &producer{
		conn:  conn,
		topic: topic,
	}
}

func (p *producer) createMessage(value string) {
	msg := newMessage(value)
	p.topic.addMessage(msg)
}

func (p *producer) produceMessages() {
	reader := bufio.NewReader(p.conn)

	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			p.conn.Close()
			p.connected = false

			return
		}

		p.createMessage(message)
	}
}
