package broker

import (
	"net"
	"time"
)

type consumer struct {
	connected  bool
	conn       net.Conn
	topic      *topic
	lastOffset int
}

func newConsumer(conn net.Conn, topic *topic) *consumer {
	return &consumer{
		conn:  conn,
		topic: topic,
	}
}

func (c *consumer) consumeMessages() {
	for {
		if c.lastOffset < len(c.topic.messages) {
			for i := c.lastOffset; i < len(c.topic.messages); i++ {
				_, err := c.conn.Write([]byte(c.topic.messages[i].value))
				if err != nil {
					c.conn.Close()
					c.connected = false

					return
				}
			}

			c.lastOffset = len(c.topic.messages)
		}

		// TODO: create channel to alert when has new message
		time.Sleep(1 * time.Second)
	}
}
