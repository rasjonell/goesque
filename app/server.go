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

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Printf("Error accepting connection: %v", err)
			continue
		}
		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {
	defer conn.Close()
	log.Println("Handling connection from", conn.RemoteAddr())

	for {
		messageSizeBuffer := make([]byte, unsafe.Sizeof(kafka.SizeMessage{}))
		conn.Read(messageSizeBuffer)
		msg := kafka.NewSizeMessageFromBuffer(messageSizeBuffer)

		requestBuffer := make([]byte, msg.Size)
		n, err := conn.Read(requestBuffer)
		if err != nil {
			log.Println("Read error:", err)
			return
		}
		if n == 0 {
			break
		}

		header := kafka.NewHeader(requestBuffer)

		response := kafka.GenerateResponse(header)
		conn.Write(response)
	}
}
