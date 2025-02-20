package script

import (
	"encoding/binary"
	"fmt"
)

func decodeBytesToInt64(data []byte) (int64, error) {
	if len(data) == 0 {
		return 0, nil
	}

	var result int64
	for i, b := range data {
		result |= int64(b) << (8 * i)
	}

	if data[len(data)-1]&0x80 != 0 {
		result &= ^(int64(0x80) << (8 * (len(data) - 1)))
		result = -result
	}

	return result, nil
}

func encodeInt64ToBytes(n int64) []byte {
	if n == 0 {
		return []byte{}
	}

	negative := n < 0
	if negative {
		n = -n
	}

	var result []byte
	for n != 0 {
		result = append(result, byte(n&0xff))
		n >>= 8
	}

	if result[len(result)-1]&0x80 != 0 {
		if negative {
			result = append(result, 0x80)
		} else {
			result = append(result, 0x00)
		}
	} else if negative {
		result[len(result)-1] |= 0x80
	}

	return result
}

func pushEncodedData(data []byte) ([]byte, error) {
	n := len(data)
	if n > MaxPushSize {
		return nil, fmt.Errorf("data push size %d exceeds maximum allowed %d", n, MaxPushSize)
	}

	var result []byte
	switch {
	case n < OP_PUSHDATA1:
		result = append(result, byte(n))
	case n <= 0xff:
		result = append(result, OP_PUSHDATA1, byte(n))
	case n <= 0xffff:
		result = append(result, OP_PUSHDATA2)
		tmp := make([]byte, 2)
		binary.LittleEndian.PutUint16(tmp, uint16(n))
		result = append(result, tmp...)
	default:
		result = append(result, OP_PUSHDATA4)
		tmp := make([]byte, 4)
		binary.LittleEndian.PutUint32(tmp, uint32(n))
		result = append(result, tmp...)
	}

	result = append(result, data...)
	return result, nil
}
