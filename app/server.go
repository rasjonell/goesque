package main

import (
	"log"
	"net"
	"unsafe"

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

	msg := &kafka.Message{}
	messageSizeBuffer := make([]byte, unsafe.Sizeof(msg))
	conn.Read(messageSizeBuffer)
	msg = kafka.NewMessage(messageSizeBuffer)

	requestBuffer := make([]byte, msg.Size)
	conn.Read(requestBuffer)
	header := kafka.NewHeader(requestBuffer)

	totalLength := len(msg.Bytes()) + len(header.Bytes())
	response := make([]byte, totalLength)

	offset := 0
	offset = copy(response[offset:], msg.Bytes())
	offset = copy(response[offset:], header.Bytes())

	conn.Write(response)
}
