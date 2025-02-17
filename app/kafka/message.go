package kafka

import (
	"encoding/binary"
	"reflect"
)

type SizeMessage struct {
	Size int32
}

func (h *SizeMessage) Bytes() []byte {
	msgType := reflect.TypeOf(*h)
	size := int(msgType.Size())

	b := make([]byte, size)
	binary.BigEndian.PutUint32(b[0:4], uint32(h.Size))
	return b
}

func NewSizeMessageFromBuffer(sizeBuffer []byte) *SizeMessage {
	return &SizeMessage{
		Size: int32(binary.BigEndian.Uint32(sizeBuffer)),
	}
}

func NewSizeMessage(size int32) *SizeMessage {
	return &SizeMessage{
		Size: size,
	}
}
