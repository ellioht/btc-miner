package core

import (
	"bytes"
	"encoding/binary"
	"github.com/ellioht/btc-miner/common"
	"github.com/ellioht/btc-miner/crypto"
)

type Header struct {
	Version       uint32
	PrevBlockHash common.Hash
	MerkleRoot    common.Hash
	Timestamp     uint32
	Bits          uint32
	Nonce         uint32
}

func NewHeader(version uint32, prevBlockHash common.Hash, merkleRoot common.Hash, timestamp uint32, bits uint32, nonce uint32) *Header {
	return &Header{
		Version:       version,
		PrevBlockHash: prevBlockHash,
		MerkleRoot:    merkleRoot,
		Timestamp:     timestamp,
		Bits:          bits,
		Nonce:         nonce,
	}
}

func (h *Header) Hash() common.Hash {
	return common.HashFromBytes(crypto.HashData(h.Serialize()))
}

func (h *Header) Serialize() []byte {
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, h.Version)
	buf.Write(h.PrevBlockHash[:])
	buf.Write(h.MerkleRoot[:])
	binary.Write(buf, binary.LittleEndian, h.Timestamp)
	binary.Write(buf, binary.LittleEndian, h.Bits)
	binary.Write(buf, binary.LittleEndian, h.Nonce)
	return buf.Bytes()
}

func (h *Header) UpdateNonce(nonce uint32) {
	h.Nonce = nonce
}
