package kafka

import (
	"log"
)

type Bytable interface {
	Bytes() []byte
}

type PlaceholderApi struct{}

func (p *PlaceholderApi) Bytes() []byte {
	return []byte{}
}

func GenerateResponse(header *HeaderV2) []byte {
	var body Bytable

	switch header.RequestApiKey {
	case API_ApiVersions:
		body = NewApiVersionsBody(header)

	default:
		log.Fatalf("FUCKY WUCKY api key: %v", header.RequestApiKey)
	}

	return buildResponse(header, body)
}

func buildResponse(header *HeaderV2, body Bytable) []byte {
	// Doing this because complete header is not allowed for this stage :(
	incompleteHeader := HeaderV0{
		CorrelationId: header.CorrelationId,
	}

	bodyBytes := body.Bytes()
	headerBytes := incompleteHeader.Bytes()

	size := len(bodyBytes) + len(headerBytes)
	sizeMsg := NewSizeMessage(int32(size))
	sizeBytes := sizeMsg.Bytes()

	totalLength := len(sizeBytes) + size
	response := make([]byte, totalLength)

	offset := 0
	copy(response[offset:], sizeBytes)
	offset += len(sizeBytes)
	copy(response[offset:], headerBytes)
	offset += len(headerBytes)
	copy(response[offset:], bodyBytes)

	return response
}
