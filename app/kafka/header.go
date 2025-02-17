package kafka

import (
	"encoding/binary"
	"log"
)

type HeaderV0 struct {
	CorrelationId int32
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

func (h *HeaderV0) Bytes() []byte {
	bytes, err := ToBytes(h)
	if err != nil {
		log.Fatalf("Error converting header to bytes: %v", err)
	}

	return bytes
}

func (h *HeaderV2) Bytes() []byte {
	bytes, err := ToBytes(h)
	if err != nil {
		log.Fatalf("Error converting header to bytes: %v", err)
	}

	return bytes
}

func NewHeader(headerBuffer []byte) *HeaderV2 {
	header := &HeaderV2{}
	header.RequestApiKey = int16(binary.BigEndian.Uint16(headerBuffer[0:2]))
	header.RequestApiVersion = int16(binary.BigEndian.Uint16(headerBuffer[2:4]))
	header.CorrelationId = int32(binary.BigEndian.Uint32(headerBuffer[4:12]))

	return header
}
