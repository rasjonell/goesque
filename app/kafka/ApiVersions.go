package kafka

import (
	"encoding/binary"
	"reflect"
)

type ApiVersionsBody struct {
	ErrorCode int16
}

func (a *ApiVersionsBody) Bytes() []byte {
	bodyType := reflect.TypeOf(*a)
	size := int(bodyType.Size())

	b := make([]byte, size)
	binary.BigEndian.PutUint16(b[0:2], uint16(a.ErrorCode))
	return b
}

func NewApiVersionsBody(_ *HeaderV2) *ApiVersionsBody {
	// For now
	return &ApiVersionsBody{
		ErrorCode: UNSUPPORTED_VERSION,
	}
}
