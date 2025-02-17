package kafka

import (
	"log"
)

type HeaderV0 struct {
	// Ignored for now
	// RequestApiKey     int16
	// RequestApiVersion int16
	CorrelationId int32
}

func (h *HeaderV0) Bytes() []byte {
	bytes, err := ToBytes(h)
	if err != nil {
		log.Fatalf("Error converting header to bytes: %v", err)
	}

	return bytes
}

func NewHeader( /*id int32*/ ) *HeaderV0 {
	return &HeaderV0{
		CorrelationId: 7, // Default value for now
	}
}
