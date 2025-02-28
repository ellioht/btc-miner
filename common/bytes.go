package common

import (
	"errors"
	"fmt"
	"strconv"
)

func ReverseBytes(b []byte) []byte {
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return b
}

func CompactToUint32(bits string) (uint32, error) {
	if len(bits) != 8 {
		return 0, errors.New("bits must be an 8-character hexadecimal string")
	}

	exponent, err := strconv.ParseUint(bits[:2], 16, 8)
	if err != nil {
		return 0, fmt.Errorf("failed to parse exponent: %v", err)
	}

	coefficient, err := strconv.ParseUint(bits[2:], 16, 24)
	if err != nil {
		return 0, fmt.Errorf("failed to parse coefficient: %v", err)
	}

	result := (exponent << 24) + coefficient

	return uint32(result), nil
}
