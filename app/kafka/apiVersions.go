package kafka

import (
	"encoding/binary"
	"reflect"
)

type ApiKey struct {
	ApiKey     int16
	MinVersion int16
	MaxVersion int16
}

type ApiVersionsBody struct {
	ThrottleTimeMs int32
	ErrorCode      int16
	ApiKeysLength  byte
	TagLine        byte
}

var supportedApiKeys []ApiKey = []ApiKey{
	{
		MinVersion: 0,
		MaxVersion: 4,
		ApiKey:     API_ApiVersions,
	},
	{
		MinVersion: 0,
		MaxVersion: 0,
		ApiKey:     API_DescribeTopicPartitions,
	},
}

func (a *ApiVersionsBody) Bytes() []byte {
	bodyType := reflect.TypeOf(*a)
	apiKeyType := reflect.TypeOf(supportedApiKeys[0])
	// ApiVersionBody(4 + 2 + 1 + 1 = 8 bytes) + (# of supported API Keys * ApiKey(2 + 2 + 2 = 6) + 1(single TagBuffer byte))
	size := int(a.ApiKeysLength)*(int(apiKeyType.Size())+1) + int(bodyType.Size())

	b := make([]byte, size)
	offset := 0
	binary.BigEndian.PutUint16(b[0:], uint16(a.ErrorCode))
	offset += 2

	b[offset] = a.ApiKeysLength + 1
	offset += 1

	if a.ApiKeysLength > 0 {
		for _, apiKey := range supportedApiKeys {
			binary.BigEndian.PutUint16(b[offset:], uint16(apiKey.ApiKey))
			offset += 2
			binary.BigEndian.PutUint16(b[offset:], uint16(apiKey.MinVersion))
			offset += 2
			binary.BigEndian.PutUint16(b[offset:], uint16(apiKey.MaxVersion))
			offset += 2
			b[offset] = 0 // TagBuffer
			offset += 1
		}
	}

	binary.BigEndian.PutUint32(b[offset:], uint32(a.ThrottleTimeMs))
	offset += 4
	b[offset] = a.TagLine

	return b
}

func NewApiVersionsBody(h *HeaderV2) *ApiVersionsBody {
	if h.RequestApiVersion < 0 || h.RequestApiVersion > 4 {
		return &ApiVersionsBody{
			ThrottleTimeMs: 0,
			ApiKeysLength:  0,
			TagLine:        0,
			ErrorCode:      ERROR_UNSUPPORTED_VERSION,
		}
	}

	return &ApiVersionsBody{
		ThrottleTimeMs: 0,
		TagLine:        0,
		ErrorCode:      ERROR_NONE,
		ApiKeysLength:  byte(len(supportedApiKeys)),
	}
}
