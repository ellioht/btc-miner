package script

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
)

const (
	MaxPushSize = 520
)

func NewCoinbaseScriptBytes(script string) ([]byte, error) {
	return hex.DecodeString(script)
}

func CoinbaseScriptToString(script []byte) string {
	return hex.EncodeToString(script)
}

// DecodeCoinbaseScript decodes a coinbase script into its height and messages.
// It will return the height, then a slice of all the pushes in the script.
func DecodeCoinbaseScript(script []byte) (int64, []Message, error) {
	if len(script) < 1 {
		return 0, nil, fmt.Errorf("script too short")
	}

	var height int64
	var pushes []Message
	offset := 0

	for pushIndex := 0; offset < len(script); pushIndex++ {
		opcode := script[offset]
		offset++

		var pushSize int

		switch {
		case opcode <= OP_DATA_75:
			pushSize = int(opcode)

		case opcode == OP_PUSHDATA1:
			if offset+1 > len(script) {
				return 0, nil, fmt.Errorf("unexpected end of script after OP_PUSHDATA1")
			}
			pushSize = int(script[offset])
			offset++

		case opcode == OP_PUSHDATA2:
			if offset+2 > len(script) {
				return 0, nil, fmt.Errorf("unexpected end of script after OP_PUSHDATA2")
			}
			pushSize = int(binary.LittleEndian.Uint16(script[offset:]))
			offset += 2

		case opcode == OP_PUSHDATA4:
			if offset+4 > len(script) {
				return 0, nil, fmt.Errorf("unexpected end of script after OP_PUSHDATA4")
			}
			pushSize = int(binary.LittleEndian.Uint32(script[offset:]))
			offset += 4

		case opcode == OP_1NEGATE:
			pushes = append(pushes, Message{0xff})
			continue
		case opcode >= OP_1 && opcode <= OP_16:
			numericVal := opcode - (OP_1 - 1)
			pushes = append(pushes, Message{numericVal})
			continue

		default:
			continue
		}

		if offset+pushSize > len(script) {
			return 0, nil, fmt.Errorf("push data exceeds script length")
		}

		pushData := script[offset : offset+pushSize]
		offset += pushSize

		if pushIndex == 0 {
			n, err := decodeBytesToInt64(pushData)
			if err != nil {
				return 0, nil, fmt.Errorf("failed to decode block height: %v", err)
			}
			if n < 0 {
				return 0, nil, fmt.Errorf("decoded block height is negative")
			}
			height = n
		} else {
			pushes = append(pushes, pushData)
		}
	}

	return height, pushes, nil
}

// EncodeCoinbaseScript creates a coinbase script with the given height, and pushData,
// By adding an extra nonce to the coinbase script, we can change the merkle root
// therefore giving us more chances to find a valid block.
func EncodeCoinbaseScript(height int64, pushes ...[]byte) ([]byte, error) {
	var script []byte

	heightBytes := encodeInt64ToBytes(height)
	heightPush, err := pushEncodedData(heightBytes)
	if err != nil {
		return nil, fmt.Errorf("failed to encode height: %v", err)
	}
	script = append(script, heightPush...)

	for _, push := range pushes {
		pushPush, err := pushEncodedData(push)
		if err != nil {
			return nil, fmt.Errorf("failed to encode push: %v", err)
		}
		script = append(script, pushPush...)
	}

	return script, nil
}
