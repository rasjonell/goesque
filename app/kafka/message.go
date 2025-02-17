package kafka

import "log"

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

func NewMessage( /*size int32*/ ) *Message {
	return &Message{Size: 0}
}
