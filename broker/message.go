package broker

import "time"

type message struct {
	key   string
	value string
	date  time.Time
}

func newMessage(key string, value string) *message {
	return &message{
		key:   key,
		value: value,
		date:  time.Now().UTC(),
	}
}
