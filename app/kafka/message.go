package kafka

import (
	"encoding/binary"
	"log"
)

type Message struct {
	Size int32
}

func (h *Message) Bytes() []byte {
	bytes, err := ToBytes(h)
	if err != nil {
		log.Fatalf("Error converting message to bytes: %v", err)
	}

	return bytes
}

func NewMessage(sizeBuffer []byte) *Message {
	return &Message{
		Size: int32(binary.BigEndian.Uint32(sizeBuffer)),
	}
}
