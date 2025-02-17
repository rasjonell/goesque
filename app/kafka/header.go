package kafka

import (
	"encoding/binary"
	"reflect"
)

type HeaderV0 struct {
	CorrelationId int32
}

func (h *HeaderV0) Bytes() []byte {
	headerType := reflect.TypeOf(*h)
	size := int(headerType.Size())

	b := make([]byte, size)
	binary.BigEndian.PutUint32(b[0:4], uint32(h.CorrelationId))
	return b
}

type NullableString struct {
	Length int16
}

type HeaderV2 struct {
	RequestApiKey     int16
	RequestApiVersion int16
	CorrelationId     int32
	// Dunno for now
	// ClientId  []byte
	// TagBuffer []byte
}

func (h *HeaderV2) Bytes() []byte {
	headerType := reflect.TypeOf(*h)
	size := int(headerType.Size())

	b := make([]byte, size)
	binary.BigEndian.PutUint16(b[0:2], uint16(h.RequestApiKey))
	binary.BigEndian.PutUint16(b[2:4], uint16(h.RequestApiVersion))
	binary.BigEndian.PutUint32(b[4:8], uint32(h.CorrelationId))
	return b
}

func NewHeader(headerBuffer []byte) *HeaderV2 {
	header := &HeaderV2{}
	header.RequestApiKey = int16(binary.BigEndian.Uint16(headerBuffer[0:2]))
	header.RequestApiVersion = int16(binary.BigEndian.Uint16(headerBuffer[2:4]))
	header.CorrelationId = int32(binary.BigEndian.Uint32(headerBuffer[4:8]))

	return header
}
