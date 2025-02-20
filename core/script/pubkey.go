package script

import (
	"bytes"
	"github.com/ellioht/btc-miner/common"
)

func PubKeyScript(addr common.Address) []byte {
	var buf bytes.Buffer
	buf.WriteByte(0x76)     // OP_DUP
	buf.WriteByte(0xa9)     // OP_HASH160
	buf.WriteByte(0x14)     // Push 20 bytes
	buf.Write(addr.Bytes()) // PubKeyHash
	buf.WriteByte(0x88)     // OP_EQUALVERIFY
	buf.WriteByte(0xac)     // OP_CHECKSIG
	return buf.Bytes()
}
