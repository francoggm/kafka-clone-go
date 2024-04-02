package main

import (
	"kafka-clone/broker"
	"kafka-clone/model"
	"kafka-clone/server"
	"log"
)

func main() {
	host := "localhost"
	port := "8081"

	ch := make(chan *model.BrokerConnection)

	broker := broker.NewBroker()
	go broker.Run(ch)

	if err := server.Run(host, port, ch); err != nil {
		log.Panic(err)
	}
}
