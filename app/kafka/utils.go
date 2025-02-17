package kafka

import (
	"bytes"
	"encoding/binary"
)

func ToBytes(data any) ([]byte, error) {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.BigEndian, data)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
