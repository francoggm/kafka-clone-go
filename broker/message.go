package broker

import "time"

type message struct {
	value string
	date  time.Time
}

func newMessage(value string) *message {
	return &message{
		value: value,
		date:  time.Now().UTC(),
	}
}
