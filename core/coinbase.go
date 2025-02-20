package core

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"github.com/ellioht/btc-miner/common"
)

func MakeCoinbaseTx(coinbaseScript, pubKeyScript []byte, value uint64) (*Tx, error) {
	var buf bytes.Buffer

	// 1. Version (1)
	binary.Write(&buf, binary.LittleEndian, uint32(1))

	// 2. Input count (one input)
	buf.WriteByte(0x01)

	// ------ Input

	// 3.1 Previous input hash (32 bytes of zeroes)
	buf.Write(common.NilHash().Bytes())
	// 3.2 Previous sequence number (0xffffffff = 4 bytes)
	binary.Write(&buf, binary.LittleEndian, uint32(0xffffffff))
	// 3.3 Input script length
	writeVarInt(&buf, uint64(len(coinbaseScript)))
	// 3.4 Input script
	buf.Write(coinbaseScript)
	// 3.5 Sequence number (0xffffffff = 4 bytes)
	binary.Write(&buf, binary.LittleEndian, uint32(0xffffffff))

	// ------

	// 4. Output count (one output)
	buf.WriteByte(0x01)

	// ------ Output

	// 5.1 Output value
	binary.Write(&buf, binary.LittleEndian, value)
	// 5.2 Output script length
	writeVarInt(&buf, uint64(len(pubKeyScript)))
	// 5.3 Output script
	buf.Write(pubKeyScript)

	// ------

	// 6. LockTime 0 (4 bytes, little-endian)
	binary.Write(&buf, binary.LittleEndian, uint32(0))

	baseTx := buf.Bytes()
	txIdHash := common.HashFromBytes(baseTx)

	tx := &Tx{
		TxId: txIdHash,
		Hex:  hex.EncodeToString(baseTx),
	}

	return tx, nil
}

func writeVarInt(buf *bytes.Buffer, val uint64) {
	if val < 0xfd {
		buf.WriteByte(byte(val))
	} else if val <= 0xffff {
		buf.WriteByte(0xfd)
		binary.Write(buf, binary.LittleEndian, uint16(val))
	} else if val <= 0xffffffff {
		buf.WriteByte(0xfe)
		binary.Write(buf, binary.LittleEndian, uint32(val))
	} else {
		buf.WriteByte(0xff)
		binary.Write(buf, binary.LittleEndian, val)
	}
}
