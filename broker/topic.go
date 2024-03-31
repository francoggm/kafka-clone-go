package broker

import "sync/atomic"

var id atomic.Uint32

type topic struct {
	id       uint32
	name     string
	messages []*message
}

func newTopic(name string) *topic {
	id.Add(1)

	return &topic{
		id:       id.Load(),
		name:     name,
		messages: []*message{},
	}
}

func (t *topic) addMessage(msg *message) {
	t.messages = append(t.messages, msg)
}
