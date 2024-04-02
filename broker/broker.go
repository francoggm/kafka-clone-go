package broker

import (
	"kafka-clone/model"
)

type Broker struct {
	topics    map[string]*topic
	producers []*producer
	consumers []*consumer
}

func NewBroker() *Broker {
	return &Broker{
		topics:    make(map[string]*topic),
		producers: make([]*producer, 0),
		consumers: make([]*consumer, 0),
	}
}

func (b *Broker) Run(ch chan *model.BrokerConnection) {
	for {
		client := <-ch

		topic, ok := b.topics[client.Topic]
		if !ok {
			topic = newTopic(client.Topic)
			b.topics[client.Topic] = topic
		}

		switch client.Mode {
		case 0:
			pr := newProducer(client.Conn, topic)
			b.producers = append(b.producers, pr)

			go pr.produceMessages()
		case 1:
			cs := newConsumer(client.Conn, topic)
			b.consumers = append(b.consumers, cs)

			go cs.consumeMessages()
		}
	}
}
