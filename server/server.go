package server

import (
	"bufio"
	"kafka-clone/model"
	"log"
	"net"
	"strconv"
	"strings"
)

func Run(host string, port string, ch chan *model.BrokerConnection) error {
	listener, err := net.Listen("tcp", host+":"+port)
	if err != nil {
		return err
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		go func(c net.Conn) {
			reader := bufio.NewReader(c)

			message, err := reader.ReadString('\n')
			if err != nil {
				log.Println(err)
				return
			}

			values := strings.Split(message, " ")
			if len(values) < 2 {
				conn.Write([]byte("Invalid values!\n"))
				return
			}

			mode, err := strconv.Atoi(values[0])
			if err != nil {
				conn.Write([]byte("Invalid values!\n"))
				return
			}

			topic := values[1]
			if topic == "" {
				conn.Write([]byte("Invalid values!\n"))
				return
			}

			ch <- &model.BrokerConnection{
				Conn:  conn,
				Mode:  mode,
				Topic: topic,
			}
		}(conn)
	}
}
