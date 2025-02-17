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

	messageSizeBuffer := make([]byte, unsafe.Sizeof(kafka.SizeMessage{}))
	conn.Read(messageSizeBuffer)
	msg := kafka.NewSizeMessageFromBuffer(messageSizeBuffer)

	requestBuffer := make([]byte, msg.Size)
	conn.Read(requestBuffer)
	header := kafka.NewHeader(requestBuffer)

	response := kafka.GenerateResponse(header)

	conn.Write(response)
}
