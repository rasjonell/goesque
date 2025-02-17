package main

import (
	"fmt"
	"log"
	"net"

	"github.com/codecrafters-io/kafka-starter-go/app/kafka"
)

func main() {
	l, err := net.Listen("tcp", "0.0.0.0:9092")
	if err != nil {
		log.Fatal("Failed to bind to port 9092")
	}

	conn, err := l.Accept()
	if err != nil {
		log.Fatal("Error accepting connection: ", err.Error())
	}

	msg := kafka.NewMessage().Bytes()
	header := kafka.NewHeader().Bytes()

	totalLength := len(msg) + len(header)
	response := make([]byte, totalLength)

	fmt.Printf("%+v\n%+v\n", msg, header)

	offset := 0
	offset += copy(response[offset:], msg)
	offset += copy(response[offset:], header)

	conn.Write(response)
}
