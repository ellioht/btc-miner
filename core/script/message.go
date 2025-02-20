package script

import (
	"encoding/binary"
	"encoding/hex"
	"unicode"
)

type Message []byte

func (m Message) String() string {
	return string(m)
}

func (m Message) Hex() string {
	return hex.EncodeToString(m)
}

func (m Message) DecodedValue() interface{} {
	isText := true
	for _, r := range string(m) {
		if !unicode.IsPrint(r) && r != '\n' && r != '\r' && r != '\t' {
			isText = false
			break
		}
	}
	if isText {
		return string(m)
	}

	switch len(m) {
	case 1:
		return uint8(m[0])
	case 2:
		return binary.LittleEndian.Uint16(m)
	case 4:
		return binary.LittleEndian.Uint32(m)
	case 8:
		return binary.LittleEndian.Uint64(m)
	}

	return hex.EncodeToString(m)
}
