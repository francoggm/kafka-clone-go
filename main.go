package main

import (
	"kafka-clone/broker"
	"kafka-clone/model"
	"log"
	"net"
)

func main() {
	host := "localhost"
	port := "8081"

	ch := make(chan *model.BrokerConnection)

	broker := broker.NewBroker()
	broker.Run(ch)

	listener, err := net.Listen("tcp", host+":"+port)
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		ch <- &model.BrokerConnection{
			Conn:  conn,
			Mode:  0,
			Topic: "",
		}
	}
}
